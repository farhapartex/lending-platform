package config

import (
	"strings"
	"testing"
	"time"
)

const (
	poolAddress       = "0x0165878A594ca255338adfa4d48449f69242Eb8F"
	vaultAddress      = "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853"
	controllerAddress = "0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6"
	managerAddress    = "0x8A791620dd6260079BF849Dc5567aDC3F2FdC318"
	lensAddress       = "0x610178dA211FEF7D417bC0e6FeD39F05609AD788"
	oracleAddress     = "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9"
)

func setRequiredChainEnv(t *testing.T) {
	t.Helper()

	t.Setenv("CHAIN_RPC_URL", "http://localhost:8545")
	t.Setenv("POOL_ADDRESS", poolAddress)
	t.Setenv("VAULT_ADDRESS", vaultAddress)
	t.Setenv("CONTROLLER_ADDRESS", controllerAddress)
	t.Setenv("LIQUIDATION_MANAGER_ADDRESS", managerAddress)
	t.Setenv("LENS_ADDRESS", lensAddress)
}

func TestLoadChainDefaults(t *testing.T) {
	setRequiredChainEnv(t)

	cfg, err := loadChain()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.ChainID != 31337 {
		t.Fatalf("expected default chain id 31337, got %d", cfg.ChainID)
	}

	if cfg.Confirmations != 3 {
		t.Fatalf("expected default confirmations 3, got %d", cfg.Confirmations)
	}

	if cfg.BatchSize != 2000 {
		t.Fatalf("expected default batch size 2000, got %d", cfg.BatchSize)
	}

	if cfg.PollInterval != 5*time.Second {
		t.Fatalf("expected default poll interval 5s, got %s", cfg.PollInterval)
	}

	if cfg.SnapshotInterval != 60*time.Second {
		t.Fatalf("expected default snapshot interval 60s, got %s", cfg.SnapshotInterval)
	}

	if cfg.RequestTimeout != 15*time.Second {
		t.Fatalf("expected default request timeout 15s, got %s", cfg.RequestTimeout)
	}

	if !cfg.IndexerEnabled {
		t.Fatal("expected the indexer to be enabled by default")
	}
}

func TestLoadChainNormalizesAddresses(t *testing.T) {
	setRequiredChainEnv(t)

	cfg, err := loadChain()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Contracts.Pool != strings.ToLower(poolAddress) {
		t.Fatalf("expected the pool address to be lowercased, got %q", cfg.Contracts.Pool)
	}

	if cfg.Contracts.Lens != strings.ToLower(lensAddress) {
		t.Fatalf("expected the lens address to be lowercased, got %q", cfg.Contracts.Lens)
	}
}

func TestLoadChainReadsOverrides(t *testing.T) {
	setRequiredChainEnv(t)

	t.Setenv("CHAIN_ID", "84532")
	t.Setenv("CHAIN_CONFIRMATIONS", "12")
	t.Setenv("INDEXER_START_BLOCK", "9000000")
	t.Setenv("INDEXER_BATCH_SIZE", "500")
	t.Setenv("INDEXER_POLL_INTERVAL", "2s")
	t.Setenv("SNAPSHOT_INTERVAL", "30s")
	t.Setenv("CHAIN_REQUEST_TIMEOUT", "20s")

	cfg, err := loadChain()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.ChainID != 84532 {
		t.Fatalf("expected chain id 84532, got %d", cfg.ChainID)
	}

	if cfg.Confirmations != 12 {
		t.Fatalf("expected confirmations 12, got %d", cfg.Confirmations)
	}

	if cfg.StartBlock != 9000000 {
		t.Fatalf("expected start block 9000000, got %d", cfg.StartBlock)
	}

	if cfg.BatchSize != 500 {
		t.Fatalf("expected batch size 500, got %d", cfg.BatchSize)
	}

	if cfg.PollInterval != 2*time.Second {
		t.Fatalf("expected poll interval 2s, got %s", cfg.PollInterval)
	}

	if cfg.SnapshotInterval != 30*time.Second {
		t.Fatalf("expected snapshot interval 30s, got %s", cfg.SnapshotInterval)
	}

	if cfg.RequestTimeout != 20*time.Second {
		t.Fatalf("expected request timeout 20s, got %s", cfg.RequestTimeout)
	}
}

func TestLoadChainRejectsBadValues(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		value    string
		contains string
	}{
		{name: "chain id not a number", key: "CHAIN_ID", value: "mainnet", contains: "CHAIN_ID"},
		{name: "chain id zero", key: "CHAIN_ID", value: "0", contains: "CHAIN_ID"},
		{name: "chain id negative", key: "CHAIN_ID", value: "-1", contains: "CHAIN_ID"},
		{name: "confirmations not a number", key: "CHAIN_CONFIRMATIONS", value: "many", contains: "CHAIN_CONFIRMATIONS"},
		{name: "confirmations negative", key: "CHAIN_CONFIRMATIONS", value: "-1", contains: "CHAIN_CONFIRMATIONS"},
		{name: "start block not a number", key: "INDEXER_START_BLOCK", value: "genesis", contains: "INDEXER_START_BLOCK"},
		{name: "batch size not a number", key: "INDEXER_BATCH_SIZE", value: "lots", contains: "INDEXER_BATCH_SIZE"},
		{name: "batch size zero", key: "INDEXER_BATCH_SIZE", value: "0", contains: "INDEXER_BATCH_SIZE"},
		{name: "indexer enabled not a bool", key: "INDEXER_ENABLED", value: "maybe", contains: "INDEXER_ENABLED"},
		{name: "poll interval not a duration", key: "INDEXER_POLL_INTERVAL", value: "soon", contains: "INDEXER_POLL_INTERVAL"},
		{name: "poll interval zero", key: "INDEXER_POLL_INTERVAL", value: "0s", contains: "INDEXER_POLL_INTERVAL"},
		{name: "snapshot interval negative", key: "SNAPSHOT_INTERVAL", value: "-5s", contains: "SNAPSHOT_INTERVAL"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			setRequiredChainEnv(t)
			t.Setenv(testCase.key, testCase.value)

			_, err := loadChain()
			if err == nil {
				t.Fatalf("expected an error for %s=%q", testCase.key, testCase.value)
			}

			if !strings.Contains(err.Error(), testCase.contains) {
				t.Fatalf("expected the error to mention %s, got %v", testCase.contains, err)
			}
		})
	}
}

func TestLoadChainRejectsBadAddresses(t *testing.T) {
	cases := []struct {
		name  string
		key   string
		value string
	}{
		{name: "not hexadecimal", key: "POOL_ADDRESS", value: "0xnothexnothexnothexnothexnothexnothexnoth"},
		{name: "too short", key: "VAULT_ADDRESS", value: "0x1234"},
		{name: "missing prefix", key: "CONTROLLER_ADDRESS", value: "0165878A594ca255338adfa4d48449f69242Eb8F"},
		{name: "zero address", key: "LENS_ADDRESS", value: "0x0000000000000000000000000000000000000000"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			setRequiredChainEnv(t)
			t.Setenv(testCase.key, testCase.value)

			_, err := loadChain()
			if err == nil {
				t.Fatalf("expected an error for %s=%q", testCase.key, testCase.value)
			}

			if !strings.Contains(err.Error(), testCase.key) {
				t.Fatalf("expected the error to name %s, got %v", testCase.key, err)
			}
		})
	}
}

func TestLoadChainRequiresRPCWhenIndexing(t *testing.T) {
	setRequiredChainEnv(t)
	t.Setenv("CHAIN_RPC_URL", "")

	_, err := loadChain()
	if err == nil {
		t.Fatal("expected an error when the rpc url is missing")
	}

	if !strings.Contains(err.Error(), "CHAIN_RPC_URL") {
		t.Fatalf("expected the error to mention CHAIN_RPC_URL, got %v", err)
	}
}

func TestLoadChainRequiresContractsWhenIndexing(t *testing.T) {
	setRequiredChainEnv(t)
	t.Setenv("LENS_ADDRESS", "")
	t.Setenv("POOL_ADDRESS", "")

	_, err := loadChain()
	if err == nil {
		t.Fatal("expected an error when contract addresses are missing")
	}

	for _, expected := range []string{"POOL_ADDRESS", "LENS_ADDRESS"} {
		if !strings.Contains(err.Error(), expected) {
			t.Fatalf("expected the error to mention %s, got %v", expected, err)
		}
	}
}

func TestLoadChainAllowsMissingSettingsWhenIndexerDisabled(t *testing.T) {
	t.Setenv("INDEXER_ENABLED", "false")

	cfg, err := loadChain()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.IndexerEnabled {
		t.Fatal("expected the indexer to be disabled")
	}

	if cfg.RPCURL != "" {
		t.Fatalf("expected an empty rpc url, got %q", cfg.RPCURL)
	}
}

func TestMissingContracts(t *testing.T) {
	cfg := ChainConfig{}

	missing := cfg.MissingContracts()
	if len(missing) != 5 {
		t.Fatalf("expected all five required contracts to be missing, got %v", missing)
	}

	cfg.Contracts = ContractAddresses{
		Pool:               poolAddress,
		Vault:              vaultAddress,
		Controller:         controllerAddress,
		LiquidationManager: managerAddress,
		Lens:               lensAddress,
	}

	if missing := cfg.MissingContracts(); len(missing) != 0 {
		t.Fatalf("expected nothing missing, got %v", missing)
	}
}

func TestIndexedAddresses(t *testing.T) {
	cfg := ChainConfig{
		Contracts: ContractAddresses{
			Pool:               "0xaaa",
			Vault:              "0xbbb",
			Controller:         "0xccc",
			LiquidationManager: "0xddd",
			Oracle:             "0xeee",
		},
	}

	addresses := cfg.IndexedAddresses()
	if len(addresses) != 5 {
		t.Fatalf("expected five indexed addresses, got %v", addresses)
	}
}

func TestIndexedAddressesSkipsEmptyAndDuplicates(t *testing.T) {
	cfg := ChainConfig{
		Contracts: ContractAddresses{
			Pool:               "0xaaa",
			Vault:              "0xaaa",
			Controller:         "",
			LiquidationManager: "0xbbb",
			Oracle:             "",
		},
	}

	addresses := cfg.IndexedAddresses()
	if len(addresses) != 2 {
		t.Fatalf("expected duplicates and blanks to be dropped, got %v", addresses)
	}

	if addresses[0] != "0xaaa" || addresses[1] != "0xbbb" {
		t.Fatalf("expected insertion order to be preserved, got %v", addresses)
	}
}

func TestLoadIncludesChainConfig(t *testing.T) {
	setRequiredChainEnv(t)
	t.Setenv("DATABASE_URL", "postgres://localhost:5432/test")

	cfg, err := Load("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Chain.ChainID != 31337 {
		t.Fatalf("expected the chain config to be loaded, got chain id %d", cfg.Chain.ChainID)
	}

	if cfg.Chain.Contracts.Pool != strings.ToLower(poolAddress) {
		t.Fatalf("expected contract addresses to be loaded, got %q", cfg.Chain.Contracts.Pool)
	}
}
