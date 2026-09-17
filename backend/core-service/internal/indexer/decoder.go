package indexer

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/farhapartex/lending-platform/core-service/internal/chain/bindings"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

type ContractSet struct {
	Pool               common.Address
	Vault              common.Address
	Controller         common.Address
	LiquidationManager common.Address
}

func (c ContractSet) Addresses() []common.Address {
	return []common.Address{c.Pool, c.Vault, c.Controller, c.LiquidationManager}
}

type LiquidationDetail struct {
	Liquidator           string
	DebtRepaid           *big.Int
	CollateralSeized     *big.Int
	BonusAmount          *big.Int
	TriggerPrice         *big.Int
	TriggerPriceDecimals int16
	Shortfall            *big.Int
}

type DecodedEvent struct {
	Type            domain.EventType
	Kind            domain.TransactionKind
	ContractAddress string
	Actor           string
	Amount          *big.Int
	HealthFactorBps *int32
	BlockNumber     uint64
	BlockHash       string
	TxHash          string
	TxIndex         uint32
	LogIndex        uint32
	Payload         map[string]any
	Liquidation     *LiquidationDetail
}

type Decoder struct {
	contracts   ContractSet
	pool        *bindings.LendingPoolFilterer
	vault       *bindings.CollateralVaultFilterer
	controller  *bindings.LendingControllerFilterer
	liquidation *bindings.LiquidationManagerFilterer
}

func NewDecoder(contracts ContractSet) (*Decoder, error) {
	pool, err := bindings.NewLendingPoolFilterer(contracts.Pool, nil)
	if err != nil {
		return nil, fmt.Errorf("lending pool events could not be prepared: %w", err)
	}

	vault, err := bindings.NewCollateralVaultFilterer(contracts.Vault, nil)
	if err != nil {
		return nil, fmt.Errorf("collateral vault events could not be prepared: %w", err)
	}

	controller, err := bindings.NewLendingControllerFilterer(contracts.Controller, nil)
	if err != nil {
		return nil, fmt.Errorf("lending controller events could not be prepared: %w", err)
	}

	manager, err := bindings.NewLiquidationManagerFilterer(contracts.LiquidationManager, nil)
	if err != nil {
		return nil, fmt.Errorf("liquidation manager events could not be prepared: %w", err)
	}

	return &Decoder{
		contracts:   contracts,
		pool:        pool,
		vault:       vault,
		controller:  controller,
		liquidation: manager,
	}, nil
}

func (d *Decoder) Contracts() ContractSet {
	return d.contracts
}

func (d *Decoder) Decode(log types.Log) (DecodedEvent, bool) {
	switch log.Address {
	case d.contracts.Pool:
		return d.decodePool(log)
	case d.contracts.Vault:
		return d.decodeVault(log)
	case d.contracts.Controller:
		return d.decodeController(log)
	case d.contracts.LiquidationManager:
		return d.decodeLiquidation(log)
	default:
		return DecodedEvent{}, false
	}
}

func (d *Decoder) decodePool(log types.Log) (DecodedEvent, bool) {
	if deposit, err := d.pool.ParseDeposit(log); err == nil {
		return base(log, domain.EventTypeDeposit, domain.TransactionKindDeposit, deposit.Lender, deposit.Assets,
			map[string]any{
				"shares":         text(deposit.Shares),
				"supply_index":   text(deposit.SupplyIndex),
				"total_supplied": text(deposit.TotalSupplied),
			}), true
	}

	if withdraw, err := d.pool.ParseWithdraw(log); err == nil {
		return base(log, domain.EventTypeWithdraw, domain.TransactionKindWithdraw, withdraw.Lender, withdraw.Assets,
			map[string]any{
				"shares":         text(withdraw.Shares),
				"supply_index":   text(withdraw.SupplyIndex),
				"total_supplied": text(withdraw.TotalSupplied),
			}), true
	}

	return DecodedEvent{}, false
}

func (d *Decoder) decodeVault(log types.Log) (DecodedEvent, bool) {
	if added, err := d.vault.ParseCollateralDeposited(log); err == nil {
		return base(log, domain.EventTypeCollateralAdded, domain.TransactionKindCollateralAdded,
			added.Borrower, added.Amount,
			map[string]any{"new_collateral": text(added.NewCollateral)}), true
	}

	if removed, err := d.vault.ParseCollateralWithdrawn(log); err == nil {
		return base(log, domain.EventTypeCollateralWithdrawn, domain.TransactionKindCollateralWithdrawn,
			removed.Borrower, removed.Amount,
			map[string]any{"new_collateral": text(removed.NewCollateral)}), true
	}

	return DecodedEvent{}, false
}

func (d *Decoder) decodeController(log types.Log) (DecodedEvent, bool) {
	if borrow, err := d.controller.ParseBorrow(log); err == nil {
		event := base(log, domain.EventTypeBorrow, domain.TransactionKindBorrow, borrow.Borrower, borrow.Amount,
			map[string]any{
				"new_debt":     text(borrow.NewDebt),
				"borrow_index": text(borrow.BorrowIndex),
			})
		event.HealthFactorBps = narrowBps(borrow.HealthFactorBps)

		return event, true
	}

	if repay, err := d.controller.ParseRepay(log); err == nil {
		event := base(log, domain.EventTypeRepay, domain.TransactionKindRepay, repay.Borrower, repay.Amount,
			map[string]any{
				"new_debt": text(repay.NewDebt),
				"payer":    strings.ToLower(repay.Payer.Hex()),
			})

		return event, true
	}

	return DecodedEvent{}, false
}

func (d *Decoder) decodeLiquidation(log types.Log) (DecodedEvent, bool) {
	executed, err := d.liquidation.ParseLiquidationExecuted(log)
	if err != nil {
		return DecodedEvent{}, false
	}

	event := base(log, domain.EventTypeLiquidation, domain.TransactionKindLiquidation,
		executed.Borrower, executed.DebtRepaid,
		map[string]any{
			"liquidator":        strings.ToLower(executed.Liquidator.Hex()),
			"collateral_seized": text(executed.CollateralSeized),
			"bonus_value":       text(executed.BonusValue),
			"collateral_price":  text(executed.CollateralPrice),
			"price_decimals":    executed.PriceDecimals,
			"shortfall":         text(executed.Shortfall),
		})

	event.HealthFactorBps = narrowBps(executed.HealthFactorBeforeBps)
	event.Liquidation = &LiquidationDetail{
		Liquidator:           strings.ToLower(executed.Liquidator.Hex()),
		DebtRepaid:           executed.DebtRepaid,
		CollateralSeized:     executed.CollateralSeized,
		BonusAmount:          executed.BonusValue,
		TriggerPrice:         executed.CollateralPrice,
		TriggerPriceDecimals: int16(executed.PriceDecimals),
		Shortfall:            executed.Shortfall,
	}

	return event, true
}

func base(
	log types.Log,
	eventType domain.EventType,
	kind domain.TransactionKind,
	actor common.Address,
	amount *big.Int,
	payload map[string]any,
) DecodedEvent {
	if payload == nil {
		payload = map[string]any{}
	}

	payload["amount"] = text(amount)
	payload["actor"] = strings.ToLower(actor.Hex())

	return DecodedEvent{
		Type:            eventType,
		Kind:            kind,
		ContractAddress: strings.ToLower(log.Address.Hex()),
		Actor:           strings.ToLower(actor.Hex()),
		Amount:          amount,
		BlockNumber:     log.BlockNumber,
		BlockHash:       log.BlockHash.Hex(),
		TxHash:          log.TxHash.Hex(),
		TxIndex:         uint32(log.TxIndex),
		LogIndex:        uint32(log.Index),
		Payload:         payload,
	}
}

func text(value *big.Int) string {
	if value == nil {
		return "0"
	}

	return value.String()
}

func narrowBps(value *big.Int) *int32 {
	if value == nil || !value.IsInt64() {
		return nil
	}

	asInt64 := value.Int64()
	if asInt64 < 0 || asInt64 > int64(^uint32(0)>>1) {
		return nil
	}

	narrowed := int32(asInt64)

	return &narrowed
}
