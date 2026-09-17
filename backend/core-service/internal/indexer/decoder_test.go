package indexer_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/farhapartex/lending-platform/core-service/internal/chain/bindings"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/indexer"
)

var (
	poolAddress       = common.HexToAddress("0x0165878A594ca255338adfa4d48449f69242Eb8F")
	vaultAddress      = common.HexToAddress("0xa513E6E4b8f2a923D98304ec87F64353C4D5C853")
	controllerAddress = common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")
	managerAddress    = common.HexToAddress("0x8A791620dd6260079BF849Dc5567aDC3F2FdC318")
	alice             = common.HexToAddress("0x90F79bf6EB2c4f870365E785982E1f101E93b906")
	liquidator        = common.HexToAddress("0x976EA74026E726554dB657fA54763abd0C3a0aa9")
)

func topicFor(address common.Address) common.Hash {
	return common.BytesToHash(common.LeftPadBytes(address.Bytes(), 32))
}

func newDecoder(t *testing.T) *indexer.Decoder {
	t.Helper()

	decoder, err := indexer.NewDecoder(indexer.ContractSet{
		Pool:               poolAddress,
		Vault:              vaultAddress,
		Controller:         controllerAddress,
		LiquidationManager: managerAddress,
	})
	if err != nil {
		t.Fatalf("building the decoder failed: %v", err)
	}

	return decoder
}

func buildLog(
	t *testing.T,
	metadata *bind.MetaData,
	address common.Address,
	eventName string,
	indexed []common.Hash,
	values ...any,
) types.Log {
	t.Helper()

	parsed, err := metadata.GetAbi()
	if err != nil {
		t.Fatalf("reading the abi failed: %v", err)
	}

	event, ok := parsed.Events[eventName]
	if !ok {
		t.Fatalf("event %s is not in the abi", eventName)
	}

	var unindexed abi.Arguments

	for _, input := range event.Inputs {
		if !input.Indexed {
			unindexed = append(unindexed, input)
		}
	}

	data, err := unindexed.Pack(values...)
	if err != nil {
		t.Fatalf("packing %s failed: %v", eventName, err)
	}

	return types.Log{
		Address:     address,
		Topics:      append([]common.Hash{event.ID}, indexed...),
		Data:        data,
		BlockNumber: 42,
		BlockHash:   common.HexToHash("0xabc"),
		TxHash:      common.HexToHash("0xdef"),
		TxIndex:     1,
		Index:       3,
	}
}

func TestDecodeReadsADepositAsALenderTransaction(t *testing.T) {
	decoder := newDecoder(t)

	log := buildLog(t, bindings.LendingPoolMetaData, poolAddress, "Deposit",
		[]common.Hash{topicFor(alice)},
		big.NewInt(1_000_000), big.NewInt(900_000), big.NewInt(1), big.NewInt(150_000_000))

	event, ok := decoder.Decode(log)
	if !ok {
		t.Fatal("expected the deposit to decode")
	}

	if event.Type != domain.EventTypeDeposit || event.Kind != domain.TransactionKindDeposit {
		t.Fatalf("expected a deposit, got %s / %s", event.Type, event.Kind)
	}

	if event.Actor != "0x90f79bf6eb2c4f870365e785982e1f101e93b906" {
		t.Fatalf("expected the lender as the actor, got %s", event.Actor)
	}

	if event.Amount.Cmp(big.NewInt(1_000_000)) != 0 {
		t.Fatalf("expected the assets as the amount, got %s", event.Amount)
	}

	if event.BlockNumber != 42 || event.LogIndex != 3 {
		t.Fatalf("expected the log position to carry through, got block %d index %d", event.BlockNumber, event.LogIndex)
	}
}

func TestDecodeKeepsTheHealthFactorFromABorrow(t *testing.T) {
	decoder := newDecoder(t)

	log := buildLog(t, bindings.LendingControllerMetaData, controllerAddress, "Borrow",
		[]common.Hash{topicFor(alice)},
		big.NewInt(1_000), big.NewInt(4_000), big.NewInt(38_666), big.NewInt(1))

	event, ok := decoder.Decode(log)
	if !ok {
		t.Fatal("expected the borrow to decode")
	}

	if event.Kind != domain.TransactionKindBorrow {
		t.Fatalf("expected a borrow, got %s", event.Kind)
	}

	if event.HealthFactorBps == nil || *event.HealthFactorBps != 38_666 {
		t.Fatalf("expected the health factor to be kept, got %v", event.HealthFactorBps)
	}
}

func TestDecodeReadsALiquidationWithItsPayout(t *testing.T) {
	decoder := newDecoder(t)

	log := buildLog(t, bindings.LiquidationManagerMetaData, managerAddress, "LiquidationExecuted",
		[]common.Hash{topicFor(alice), topicFor(liquidator)},
		big.NewInt(5_100), big.NewInt(1_800), big.NewInt(255), big.NewInt(9_098),
		big.NewInt(290_000_000_000), uint8(8), big.NewInt(0))

	event, ok := decoder.Decode(log)
	if !ok {
		t.Fatal("expected the liquidation to decode")
	}

	if event.Liquidation == nil {
		t.Fatal("expected the liquidation detail to be filled in")
	}

	if event.Actor != "0x90f79bf6eb2c4f870365e785982e1f101e93b906" {
		t.Fatalf("expected the borrower as the actor, got %s", event.Actor)
	}

	if event.Liquidation.Liquidator != "0x976ea74026e726554db657fa54763abd0c3a0aa9" {
		t.Fatalf("expected the liquidator to be recorded, got %s", event.Liquidation.Liquidator)
	}

	if event.Liquidation.CollateralSeized.Cmp(big.NewInt(1_800)) != 0 {
		t.Fatalf("expected the seized collateral, got %s", event.Liquidation.CollateralSeized)
	}

	if event.Liquidation.TriggerPriceDecimals != 8 {
		t.Fatalf("expected the price decimals, got %d", event.Liquidation.TriggerPriceDecimals)
	}

	if event.HealthFactorBps == nil || *event.HealthFactorBps != 9_098 {
		t.Fatalf("expected the health factor before, got %v", event.HealthFactorBps)
	}
}

func TestDecodeIgnoresLogsFromOtherContracts(t *testing.T) {
	decoder := newDecoder(t)

	log := buildLog(t, bindings.LendingPoolMetaData, common.HexToAddress("0xdead"), "Deposit",
		[]common.Hash{topicFor(alice)},
		big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(1))

	if _, ok := decoder.Decode(log); ok {
		t.Fatal("expected a log from an unknown contract to be ignored")
	}
}

func TestDecodeIgnoresEventsItDoesNotTrack(t *testing.T) {
	decoder := newDecoder(t)

	log := buildLog(t, bindings.LendingPoolMetaData, poolAddress, "MinDepositChanged",
		nil, big.NewInt(0), big.NewInt(1_000_000))

	if _, ok := decoder.Decode(log); ok {
		t.Fatal("expected an untracked event to be ignored")
	}
}
