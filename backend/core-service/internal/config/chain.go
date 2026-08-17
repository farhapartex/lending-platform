package config

import (
	"fmt"
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/ethaddr"
)

type ContractAddresses struct {
	Pool               string
	Vault              string
	Controller         string
	LiquidationManager string
	Lens               string
	Oracle             string
	RateModel          string
	CollateralToken    string
	DebtToken          string
}

type ChainConfig struct {
	ChainID          int64
	RPCURL           string
	Confirmations    uint64
	StartBlock       uint64
	BatchSize        uint64
	PollInterval     time.Duration
	SnapshotInterval time.Duration
	RequestTimeout   time.Duration
	IndexerEnabled   bool
	Contracts        ContractAddresses
}

func loadChain() (ChainConfig, error) {
	cfg := ChainConfig{
		RPCURL: env("CHAIN_RPC_URL", ""),
	}

	chainID, err := envInt64("CHAIN_ID", 31337)
	if err != nil {
		return ChainConfig{}, err
	}
	cfg.ChainID = chainID

	confirmations, err := envUint64("CHAIN_CONFIRMATIONS", 3)
	if err != nil {
		return ChainConfig{}, err
	}
	cfg.Confirmations = confirmations

	startBlock, err := envUint64("INDEXER_START_BLOCK", 0)
	if err != nil {
		return ChainConfig{}, err
	}
	cfg.StartBlock = startBlock

	batchSize, err := envUint64("INDEXER_BATCH_SIZE", 2000)
	if err != nil {
		return ChainConfig{}, err
	}
	cfg.BatchSize = batchSize

	indexerEnabled, err := envBool("INDEXER_ENABLED", true)
	if err != nil {
		return ChainConfig{}, err
	}
	cfg.IndexerEnabled = indexerEnabled

	durations := []struct {
		key      string
		fallback time.Duration
		target   *time.Duration
	}{
		{"INDEXER_POLL_INTERVAL", 5 * time.Second, &cfg.PollInterval},
		{"SNAPSHOT_INTERVAL", 60 * time.Second, &cfg.SnapshotInterval},
		{"CHAIN_REQUEST_TIMEOUT", 15 * time.Second, &cfg.RequestTimeout},
	}

	for _, duration := range durations {
		value, err := envDuration(duration.key, duration.fallback)
		if err != nil {
			return ChainConfig{}, err
		}
		*duration.target = value
	}

	contracts, err := loadContractAddresses()
	if err != nil {
		return ChainConfig{}, err
	}
	cfg.Contracts = contracts

	if err := cfg.validate(); err != nil {
		return ChainConfig{}, err
	}

	return cfg, nil
}

func loadContractAddresses() (ContractAddresses, error) {
	addresses := ContractAddresses{}

	fields := []struct {
		key    string
		target *string
	}{
		{"POOL_ADDRESS", &addresses.Pool},
		{"VAULT_ADDRESS", &addresses.Vault},
		{"CONTROLLER_ADDRESS", &addresses.Controller},
		{"LIQUIDATION_MANAGER_ADDRESS", &addresses.LiquidationManager},
		{"LENS_ADDRESS", &addresses.Lens},
		{"ORACLE_ADDRESS", &addresses.Oracle},
		{"RATE_MODEL_ADDRESS", &addresses.RateModel},
		{"COLLATERAL_TOKEN_ADDRESS", &addresses.CollateralToken},
		{"DEBT_TOKEN_ADDRESS", &addresses.DebtToken},
	}

	for _, field := range fields {
		raw := env(field.key, "")
		if raw == "" {
			continue
		}

		normalized, err := ethaddr.NormalizeNonZero(raw)
		if err != nil {
			return ContractAddresses{}, fmt.Errorf("%s is not a usable contract address: %w", field.key, err)
		}

		*field.target = normalized
	}

	return addresses, nil
}

func (c ChainConfig) validate() error {
	if c.ChainID < 1 {
		return fmt.Errorf("CHAIN_ID must be a positive integer: got %d", c.ChainID)
	}

	if c.BatchSize < 1 {
		return fmt.Errorf("INDEXER_BATCH_SIZE must be at least 1: got %d", c.BatchSize)
	}

	if !c.IndexerEnabled {
		return nil
	}

	if c.RPCURL == "" {
		return fmt.Errorf("CHAIN_RPC_URL must be set when the indexer is enabled")
	}

	missing := c.MissingContracts()
	if len(missing) > 0 {
		return fmt.Errorf("these contract addresses must be set when the indexer is enabled: %v", missing)
	}

	return nil
}

func (c ChainConfig) MissingContracts() []string {
	required := []struct {
		key   string
		value string
	}{
		{"POOL_ADDRESS", c.Contracts.Pool},
		{"VAULT_ADDRESS", c.Contracts.Vault},
		{"CONTROLLER_ADDRESS", c.Contracts.Controller},
		{"LIQUIDATION_MANAGER_ADDRESS", c.Contracts.LiquidationManager},
		{"LENS_ADDRESS", c.Contracts.Lens},
	}

	missing := make([]string, 0, len(required))

	for _, field := range required {
		if field.value == "" {
			missing = append(missing, field.key)
		}
	}

	return missing
}

func (c ChainConfig) IndexedAddresses() []string {
	candidates := []string{
		c.Contracts.Pool,
		c.Contracts.Vault,
		c.Contracts.Controller,
		c.Contracts.LiquidationManager,
		c.Contracts.Oracle,
	}

	addresses := make([]string, 0, len(candidates))
	seen := make(map[string]struct{}, len(candidates))

	for _, candidate := range candidates {
		if candidate == "" {
			continue
		}

		if _, exists := seen[candidate]; exists {
			continue
		}

		seen[candidate] = struct{}{}
		addresses = append(addresses, candidate)
	}

	return addresses
}
