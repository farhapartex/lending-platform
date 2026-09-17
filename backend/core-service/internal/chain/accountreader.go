package chain

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/farhapartex/lending-platform/core-service/internal/chain/bindings"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

type AccountView struct {
	SupplyShares     *big.Int
	SupplyAssets     *big.Int
	CollateralAmount *big.Int
	CollateralValue  *big.Int
	DebtAmount       *big.Int
	DebtValue        *big.Int
	HealthFactorBps  *big.Int
	CollateralPrice  *big.Int
	IsLiquidatable   bool
	PriceStale       bool
}

type AccountReader struct {
	lens           *bindings.PositionLens
	requestTimeout time.Duration
}

func NewAccountReader(client *ethclient.Client, params MarketReaderParams) (*AccountReader, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: no chain client was supplied", domain.ErrChainUnreachable)
	}

	if !common.IsHexAddress(params.LensAddress) {
		return nil, fmt.Errorf("%w: lens address %q", domain.ErrInvalidInput, params.LensAddress)
	}

	lens, err := bindings.NewPositionLens(common.HexToAddress(params.LensAddress), client)
	if err != nil {
		return nil, fmt.Errorf("%w: binding the position lens failed: %s", domain.ErrChainUnreachable, err)
	}

	timeout := params.RequestTimeout
	if timeout <= 0 {
		timeout = defaultRequestTimeout
	}

	return &AccountReader{lens: lens, requestTimeout: timeout}, nil
}

func (r *AccountReader) ReadAccount(ctx context.Context, address string) (AccountView, error) {
	if !common.IsHexAddress(address) {
		return AccountView{}, fmt.Errorf("%w: account address %q", domain.ErrInvalidInput, address)
	}

	callCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()

	raw, err := r.lens.AccountData(&bind.CallOpts{Context: callCtx}, common.HexToAddress(address))
	if err != nil {
		return AccountView{}, fmt.Errorf(
			"%w: reading account %s failed: %s", domain.ErrChainUnreachable, address, err,
		)
	}

	return AccountView{
		SupplyShares:     raw.SupplyShares,
		SupplyAssets:     raw.SupplyAssets,
		CollateralAmount: raw.CollateralAmount,
		CollateralValue:  raw.CollateralValue,
		DebtAmount:       raw.DebtAmount,
		DebtValue:        raw.DebtValue,
		HealthFactorBps:  raw.HealthFactorBps,
		CollateralPrice:  raw.CollateralPrice,
		IsLiquidatable:   raw.IsLiquidatable,
		PriceStale:       raw.PriceStale,
	}, nil
}
