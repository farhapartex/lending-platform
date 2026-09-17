package chain

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/farhapartex/lending-platform/core-service/internal/chain/bindings"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

type TokenInfo struct {
	Address  string
	Symbol   string
	Name     string
	Decimals int16
}

func ReadToken(ctx context.Context, client *ethclient.Client, address string, timeout time.Duration) (TokenInfo, error) {
	if client == nil {
		return TokenInfo{}, fmt.Errorf("%w: no chain client was supplied", domain.ErrChainUnreachable)
	}

	if !common.IsHexAddress(address) {
		return TokenInfo{}, fmt.Errorf("%w: token address %q", domain.ErrInvalidInput, address)
	}

	if timeout <= 0 {
		timeout = defaultRequestTimeout
	}

	callCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	token, err := bindings.NewMockERC20(common.HexToAddress(address), client)
	if err != nil {
		return TokenInfo{}, fmt.Errorf("%w: binding token %s failed: %s", domain.ErrChainUnreachable, address, err)
	}

	opts := &bind.CallOpts{Context: callCtx}

	symbol, err := token.Symbol(opts)
	if err != nil {
		return TokenInfo{}, fmt.Errorf("%w: reading the symbol of %s failed: %s", domain.ErrChainUnreachable, address, err)
	}

	name, err := token.Name(opts)
	if err != nil {
		return TokenInfo{}, fmt.Errorf("%w: reading the name of %s failed: %s", domain.ErrChainUnreachable, address, err)
	}

	decimals, err := token.Decimals(opts)
	if err != nil {
		return TokenInfo{}, fmt.Errorf("%w: reading the decimals of %s failed: %s", domain.ErrChainUnreachable, address, err)
	}

	return TokenInfo{
		Address:  address,
		Symbol:   symbol,
		Name:     name,
		Decimals: int16(decimals),
	}, nil
}
