package chain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

const defaultRequestTimeout = 15 * time.Second

type ClientParams struct {
	RPCURL          string
	ExpectedChainID int64
	RequestTimeout  time.Duration
}

type Client struct {
	eth            *ethclient.Client
	chainID        int64
	requestTimeout time.Duration
}

func Dial(ctx context.Context, params ClientParams) (*Client, error) {
	if params.RPCURL == "" {
		return nil, fmt.Errorf("%w: no rpc url was configured", domain.ErrChainUnreachable)
	}

	timeout := params.RequestTimeout
	if timeout <= 0 {
		timeout = defaultRequestTimeout
	}

	dialCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ethClient, err := ethclient.DialContext(dialCtx, params.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", domain.ErrChainUnreachable, err)
	}

	client := &Client{eth: ethClient, requestTimeout: timeout}

	reportedChainID, err := client.fetchChainID(ctx)
	if err != nil {
		ethClient.Close()

		return nil, err
	}

	if params.ExpectedChainID > 0 && reportedChainID != params.ExpectedChainID {
		ethClient.Close()

		return nil, fmt.Errorf(
			"%w: node reports %d but configuration expects %d",
			domain.ErrChainIDMismatch, reportedChainID, params.ExpectedChainID,
		)
	}

	client.chainID = reportedChainID

	return client, nil
}

func (c *Client) Close() {
	c.eth.Close()
}

func (c *Client) ChainID(ctx context.Context) (int64, error) {
	if c.chainID > 0 {
		return c.chainID, nil
	}

	return c.fetchChainID(ctx)
}

func (c *Client) HeadBlock(ctx context.Context) (uint64, error) {
	callCtx, cancel := c.withTimeout(ctx)
	defer cancel()

	number, err := c.eth.BlockNumber(callCtx)
	if err != nil {
		return 0, fmt.Errorf("%w: reading the head block failed: %s", domain.ErrChainUnreachable, err)
	}

	return number, nil
}

func (c *Client) BlockByNumber(ctx context.Context, number uint64) (domain.BlockRef, error) {
	callCtx, cancel := c.withTimeout(ctx)
	defer cancel()

	header, err := c.eth.HeaderByNumber(callCtx, new(big.Int).SetUint64(number))
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			return domain.BlockRef{}, fmt.Errorf("%w: block %d", domain.ErrBlockNotFound, number)
		}

		return domain.BlockRef{}, fmt.Errorf("%w: reading block %d failed: %s", domain.ErrChainUnreachable, number, err)
	}

	return domain.BlockRef{
		Number: header.Number.Uint64(),
		Hash:   header.Hash().Hex(),
		Time:   time.Unix(int64(header.Time), 0).UTC(),
	}, nil
}

func (c *Client) SafeHeadBlock(ctx context.Context, confirmations uint64) (uint64, error) {
	head, err := c.HeadBlock(ctx)
	if err != nil {
		return 0, err
	}

	if head < confirmations {
		return 0, nil
	}

	return head - confirmations, nil
}

func (c *Client) fetchChainID(ctx context.Context) (int64, error) {
	callCtx, cancel := c.withTimeout(ctx)
	defer cancel()

	chainID, err := c.eth.ChainID(callCtx)
	if err != nil {
		return 0, fmt.Errorf("%w: reading the chain id failed: %s", domain.ErrChainUnreachable, err)
	}

	if !chainID.IsInt64() {
		return 0, fmt.Errorf("%w: node reported an unusable chain id %s", domain.ErrChainIDMismatch, chainID)
	}

	return chainID.Int64(), nil
}

func (c *Client) withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, c.requestTimeout)
}
