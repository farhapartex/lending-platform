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
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

type MarketReaderParams struct {
	LensAddress    string
	RequestTimeout time.Duration
}

type MarketReader struct {
	lens           *bindings.PositionLens
	requestTimeout time.Duration
}

func NewMarketReader(client *ethclient.Client, params MarketReaderParams) (*MarketReader, error) {
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

	return &MarketReader{lens: lens, requestTimeout: timeout}, nil
}

func (r *MarketReader) ReadMarketView(ctx context.Context) (domain.MarketView, error) {
	callCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()

	raw, err := r.lens.MarketData(&bind.CallOpts{Context: callCtx})
	if err != nil {
		return domain.MarketView{}, fmt.Errorf("%w: reading market data failed: %s", domain.ErrChainUnreachable, err)
	}

	return marketViewFrom(raw)
}

type bpsNarrower struct {
	err error
}

func (n *bpsNarrower) take(name string, value *big.Int) int32 {
	if n.err != nil {
		return 0
	}

	narrowed, err := bigmath.ToInt32(value)
	if err != nil {
		n.err = fmt.Errorf("%w: %s %s", domain.ErrInvalidInput, name, err)

		return 0
	}

	return narrowed
}

func marketViewFrom(raw bindings.MarketData) (domain.MarketView, error) {
	narrower := &bpsNarrower{}

	view := domain.MarketView{
		TotalSupplied:           raw.TotalSupplied,
		TotalBorrowed:           raw.TotalBorrowed,
		AvailableLiquidity:      raw.AvailableLiquidity,
		UtilizationBps:          narrower.take("utilizationBps", raw.UtilizationBps),
		SupplyRatePerSecond:     raw.SupplyRatePerSecond,
		BorrowRatePerSecond:     raw.BorrowRatePerSecond,
		SupplyAprBps:            narrower.take("supplyAprBps", raw.SupplyAprBps),
		BorrowAprBps:            narrower.take("borrowAprBps", raw.BorrowAprBps),
		SupplyIndex:             raw.SupplyIndex,
		BorrowIndex:             raw.BorrowIndex,
		MaxLtvBps:               narrower.take("maxLtvBps", raw.MaxLtvBps),
		LiquidationThresholdBps: narrower.take("liquidationThresholdBps", raw.LiquidationThresholdBps),
		LiquidationBonusBps:     narrower.take("liquidationBonusBps", raw.LiquidationBonusBps),
		KinkUtilizationBps:      narrower.take("kinkUtilizationBps", raw.KinkUtilizationBps),
		ReserveFactorBps:        narrower.take("reserveFactorBps", raw.ReserveFactorBps),
		MinDeposit:              raw.MinDeposit,
		AccruedReserves:         raw.AccruedReserves,
		DepositsPaused:          raw.DepositsPaused,
		BorrowPaused:            raw.BorrowPaused,
	}

	if narrower.err != nil {
		return domain.MarketView{}, narrower.err
	}

	return view, nil
}
