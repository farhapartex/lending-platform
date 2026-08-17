package chain_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/chain"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

type rpcRequest struct {
	ID     json.RawMessage `json:"id"`
	Method string          `json:"method"`
	Params []any           `json:"params"`
}

type fakeNode struct {
	server    *httptest.Server
	responses map[string]any
	failWith  map[string]string
	calls     map[string]int
	delay     time.Duration
}

func newFakeNode(t *testing.T) *fakeNode {
	t.Helper()

	node := &fakeNode{
		responses: map[string]any{
			"eth_chainId":          "0x7a69",
			"eth_blockNumber":      "0x1b4",
			"eth_getBlockByNumber": defaultHeader(),
		},
		failWith: map[string]string{},
		calls:    map[string]int{},
	}

	node.server = httptest.NewServer(http.HandlerFunc(node.handle))
	t.Cleanup(node.server.Close)

	return node
}

func defaultHeader() map[string]any {
	return map[string]any{
		"number":           "0x1b4",
		"hash":             "0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd",
		"parentHash":       "0x0000000000000000000000000000000000000000000000000000000000000000",
		"sha3Uncles":       "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
		"logsBloom":        "0x" + repeatZero(512),
		"transactionsRoot": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"stateRoot":        "0x0000000000000000000000000000000000000000000000000000000000000000",
		"receiptsRoot":     "0x0000000000000000000000000000000000000000000000000000000000000000",
		"miner":            "0x0000000000000000000000000000000000000000",
		"difficulty":       "0x0",
		"extraData":        "0x",
		"size":             "0x220",
		"gasLimit":         "0x1c9c380",
		"gasUsed":          "0x0",
		"timestamp":        "0x65b7f000",
		"nonce":            "0x0000000000000000",
		"baseFeePerGas":    "0x7",
		"mixHash":          "0x0000000000000000000000000000000000000000000000000000000000000000",
		"uncles":           []any{},
		"transactions":     []any{},
	}
}

func repeatZero(count int) string {
	out := make([]byte, count)
	for index := range out {
		out[index] = '0'
	}

	return string(out)
}

func (n *fakeNode) handle(w http.ResponseWriter, r *http.Request) {
	if n.delay > 0 {
		time.Sleep(n.delay)
	}

	var request rpcRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	n.calls[request.Method]++

	w.Header().Set("Content-Type", "application/json")

	if message, failing := n.failWith[request.Method]; failing {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"jsonrpc": "2.0",
			"id":      request.ID,
			"error":   map[string]any{"code": -32000, "message": message},
		})

		return
	}

	_ = json.NewEncoder(w).Encode(map[string]any{
		"jsonrpc": "2.0",
		"id":      request.ID,
		"result":  n.responses[request.Method],
	})
}

func (n *fakeNode) url() string {
	return n.server.URL
}

func dialFake(t *testing.T, node *fakeNode, expectedChainID int64) *chain.Client {
	t.Helper()

	client, err := chain.Dial(context.Background(), chain.ClientParams{
		RPCURL:          node.url(),
		ExpectedChainID: expectedChainID,
		RequestTimeout:  2 * time.Second,
	})
	if err != nil {
		t.Fatalf("unexpected dial error: %v", err)
	}

	t.Cleanup(client.Close)

	return client
}

func TestDialRejectsEmptyURL(t *testing.T) {
	_, err := chain.Dial(context.Background(), chain.ClientParams{})

	if !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected ErrChainUnreachable, got %v", err)
	}
}

func TestDialRejectsMalformedURL(t *testing.T) {
	_, err := chain.Dial(context.Background(), chain.ClientParams{
		RPCURL:          "ftp://example.com",
		ExpectedChainID: 31337,
		RequestTimeout:  time.Second,
	})

	if !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected ErrChainUnreachable, got %v", err)
	}
}

func TestDialRejectsChainIDBeyondInt64(t *testing.T) {
	node := newFakeNode(t)
	node.responses["eth_chainId"] = "0xffffffffffffffffff"

	_, err := chain.Dial(context.Background(), chain.ClientParams{
		RPCURL:          node.url(),
		ExpectedChainID: 0,
		RequestTimeout:  time.Second,
	})

	if !errors.Is(err, domain.ErrChainIDMismatch) {
		t.Fatalf("expected ErrChainIDMismatch, got %v", err)
	}
}

func TestChainIDRefetchesWhenNodeReportsZero(t *testing.T) {
	node := newFakeNode(t)
	node.responses["eth_chainId"] = "0x0"

	client := dialFake(t, node, 0)

	callsAfterDial := node.calls["eth_chainId"]

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if chainID != 0 {
		t.Fatalf("expected the reported chain id 0, got %d", chainID)
	}

	if node.calls["eth_chainId"] <= callsAfterDial {
		t.Fatal("expected an uncached chain id to be fetched again")
	}
}

func TestBlockByNumberRejectsIncompleteHeader(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	header := defaultHeader()
	delete(header, "number")
	node.responses["eth_getBlockByNumber"] = header

	_, err := client.BlockByNumber(context.Background(), 436)
	if !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected an incomplete header to be reported as unreachable, got %v", err)
	}

	if !strings.Contains(err.Error(), "number") {
		t.Fatalf("expected the error to name the missing field, got %v", err)
	}
}

func TestDialRejectsUnreachableNode(t *testing.T) {
	_, err := chain.Dial(context.Background(), chain.ClientParams{
		RPCURL:          "http://127.0.0.1:1",
		ExpectedChainID: 31337,
		RequestTimeout:  time.Second,
	})

	if !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected ErrChainUnreachable, got %v", err)
	}
}

func TestDialRejectsMismatchedChainID(t *testing.T) {
	node := newFakeNode(t)

	_, err := chain.Dial(context.Background(), chain.ClientParams{
		RPCURL:          node.url(),
		ExpectedChainID: 1,
		RequestTimeout:  time.Second,
	})

	if !errors.Is(err, domain.ErrChainIDMismatch) {
		t.Fatalf("expected ErrChainIDMismatch, got %v", err)
	}
}

func TestDialAcceptsMatchingChainID(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if chainID != 31337 {
		t.Fatalf("expected chain id 31337, got %d", chainID)
	}
}

func TestDialSkipsVerificationWhenNoChainIDExpected(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 0)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if chainID != 31337 {
		t.Fatalf("expected the reported chain id, got %d", chainID)
	}
}

func TestDialFailsWhenChainIDCallFails(t *testing.T) {
	node := newFakeNode(t)
	node.failWith["eth_chainId"] = "method not supported"

	_, err := chain.Dial(context.Background(), chain.ClientParams{
		RPCURL:          node.url(),
		ExpectedChainID: 31337,
		RequestTimeout:  time.Second,
	})

	if !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected ErrChainUnreachable, got %v", err)
	}
}

func TestDialUsesDefaultTimeoutWhenUnset(t *testing.T) {
	node := newFakeNode(t)

	client, err := chain.Dial(context.Background(), chain.ClientParams{
		RPCURL:          node.url(),
		ExpectedChainID: 31337,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Cleanup(client.Close)

	if _, err := client.HeadBlock(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestChainIDIsCachedAfterDial(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	callsAfterDial := node.calls["eth_chainId"]

	for range 3 {
		if _, err := client.ChainID(context.Background()); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if node.calls["eth_chainId"] != callsAfterDial {
		t.Fatalf("expected the chain id to be cached, saw %d calls", node.calls["eth_chainId"])
	}
}

func TestHeadBlock(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	head, err := client.HeadBlock(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if head != 436 {
		t.Fatalf("expected block 436, got %d", head)
	}
}

func TestHeadBlockPropagatesFailure(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	node.failWith["eth_blockNumber"] = "node is syncing"

	_, err := client.HeadBlock(context.Background())
	if !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected ErrChainUnreachable, got %v", err)
	}
}

func TestBlockByNumber(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	block, err := client.BlockByNumber(context.Background(), 436)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if block.Number != 436 {
		t.Fatalf("expected block number 436, got %d", block.Number)
	}

	if block.Hash == "" {
		t.Fatal("expected a block hash")
	}

	if block.Time.UTC() != time.Unix(0x65b7f000, 0).UTC() {
		t.Fatalf("expected the header timestamp, got %s", block.Time)
	}

	if block.IsZero() {
		t.Fatal("expected a populated block reference")
	}
}

func TestBlockByNumberNotFound(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	node.responses["eth_getBlockByNumber"] = nil

	_, err := client.BlockByNumber(context.Background(), 999999)
	if !errors.Is(err, domain.ErrBlockNotFound) {
		t.Fatalf("expected ErrBlockNotFound, got %v", err)
	}
}

func TestBlockByNumberPropagatesFailure(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	node.failWith["eth_getBlockByNumber"] = "database corrupted"

	_, err := client.BlockByNumber(context.Background(), 436)
	if !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected ErrChainUnreachable, got %v", err)
	}
}

func TestSafeHeadBlock(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	cases := []struct {
		name          string
		confirmations uint64
		want          uint64
	}{
		{name: "no confirmations", confirmations: 0, want: 436},
		{name: "three confirmations", confirmations: 3, want: 433},
		{name: "exactly at head", confirmations: 436, want: 0},
		{name: "more than head", confirmations: 500, want: 0},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := client.SafeHeadBlock(context.Background(), testCase.confirmations)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != testCase.want {
				t.Fatalf("expected %d, got %d", testCase.want, got)
			}
		})
	}
}

func TestSafeHeadBlockPropagatesFailure(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	node.failWith["eth_blockNumber"] = "unavailable"

	_, err := client.SafeHeadBlock(context.Background(), 3)
	if !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected ErrChainUnreachable, got %v", err)
	}
}

func TestRequestTimeoutIsApplied(t *testing.T) {
	node := newFakeNode(t)
	node.delay = 300 * time.Millisecond

	slowClient, err := chain.Dial(context.Background(), chain.ClientParams{
		RPCURL:          node.url(),
		ExpectedChainID: 31337,
		RequestTimeout:  50 * time.Millisecond,
	})
	if err == nil {
		t.Cleanup(slowClient.Close)
		t.Fatal("expected the dial to time out while reading the chain id")
	}

	if !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected ErrChainUnreachable, got %v", err)
	}
}

func TestCancelledContextIsReported(t *testing.T) {
	node := newFakeNode(t)
	client := dialFake(t, node, 31337)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := client.HeadBlock(ctx); err == nil {
		t.Fatal("expected an error for a cancelled context")
	}
}

func TestBlockRefIsZero(t *testing.T) {
	if !(domain.BlockRef{}).IsZero() {
		t.Fatal("expected an empty block reference to report zero")
	}

	if (domain.BlockRef{Number: 1}).IsZero() {
		t.Fatal("expected a numbered block reference not to report zero")
	}

	if (domain.BlockRef{Hash: "0xabc"}).IsZero() {
		t.Fatal("expected a hashed block reference not to report zero")
	}
}
