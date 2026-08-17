package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	EnvLocal      = "local"
	EnvStaging    = "staging"
	EnvProduction = "production"

	localIDMaskSecret     = "local-development-id-mask-secret"
	minIDMaskSecretLength = 32
)

func (c Config) EffectiveIDMaskSecret() string {
	if c.IDMaskSecret != "" {
		return c.IDMaskSecret
	}

	return localIDMaskSecret
}

type Config struct {
	AppEnv         string
	ServiceName    string
	ServiceVersion string
	HTTPPort       int
	LogLevel       string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	ShutdownGrace  time.Duration

	DatabaseURL             string
	DatabaseMaxOpenConns    int
	DatabaseMaxIdleConns    int
	DatabaseConnMaxLifetime time.Duration
	DatabaseConnMaxIdleTime time.Duration
	DatabaseLogQueries      bool

	IDMaskSecret string

	Chain ChainConfig
}

func (c Config) IsLocal() bool {
	return c.AppEnv == EnvLocal
}

func Load(serviceVersion string) (Config, error) {
	cfg := Config{
		AppEnv:         env("APP_ENV", EnvLocal),
		ServiceName:    env("SERVICE_NAME", "core-service"),
		ServiceVersion: serviceVersion,
		LogLevel:       env("LOG_LEVEL", "info"),
		DatabaseURL:    env("DATABASE_URL", ""),
		IDMaskSecret:   env("ID_MASK_SECRET", ""),
	}

	port, err := envInt("HTTP_PORT", 8080)
	if err != nil {
		return Config{}, err
	}
	cfg.HTTPPort = port

	maxOpen, err := envInt("DATABASE_MAX_OPEN_CONNS", 20)
	if err != nil {
		return Config{}, err
	}
	cfg.DatabaseMaxOpenConns = maxOpen

	maxIdle, err := envInt("DATABASE_MAX_IDLE_CONNS", 10)
	if err != nil {
		return Config{}, err
	}
	cfg.DatabaseMaxIdleConns = maxIdle

	logQueries, err := envBool("DATABASE_LOG_QUERIES", false)
	if err != nil {
		return Config{}, err
	}
	cfg.DatabaseLogQueries = logQueries

	durations := []struct {
		key    string
		fallb  time.Duration
		target *time.Duration
	}{
		{"HTTP_READ_TIMEOUT", 10 * time.Second, &cfg.ReadTimeout},
		{"HTTP_WRITE_TIMEOUT", 15 * time.Second, &cfg.WriteTimeout},
		{"HTTP_IDLE_TIMEOUT", 60 * time.Second, &cfg.IdleTimeout},
		{"SHUTDOWN_GRACE_PERIOD", 15 * time.Second, &cfg.ShutdownGrace},
		{"DATABASE_CONN_MAX_LIFETIME", 30 * time.Minute, &cfg.DatabaseConnMaxLifetime},
		{"DATABASE_CONN_MAX_IDLE_TIME", 5 * time.Minute, &cfg.DatabaseConnMaxIdleTime},
	}

	for _, d := range durations {
		value, err := envDuration(d.key, d.fallb)
		if err != nil {
			return Config{}, err
		}
		*d.target = value
	}

	chain, err := loadChain()
	if err != nil {
		return Config{}, err
	}
	cfg.Chain = chain

	if err := cfg.validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) validate() error {
	switch c.AppEnv {
	case EnvLocal, EnvStaging, EnvProduction:
	default:
		return fmt.Errorf("APP_ENV must be one of %s, %s, %s: got %q", EnvLocal, EnvStaging, EnvProduction, c.AppEnv)
	}

	switch strings.ToLower(c.LogLevel) {
	case "debug", "info", "warn", "error":
	default:
		return fmt.Errorf("LOG_LEVEL must be one of debug, info, warn, error: got %q", c.LogLevel)
	}

	if c.HTTPPort < 1 || c.HTTPPort > 65535 {
		return fmt.Errorf("HTTP_PORT must be between 1 and 65535: got %d", c.HTTPPort)
	}

	if c.ServiceName == "" {
		return fmt.Errorf("SERVICE_NAME must not be empty")
	}

	if c.DatabaseMaxOpenConns < 1 {
		return fmt.Errorf("DATABASE_MAX_OPEN_CONNS must be at least 1: got %d", c.DatabaseMaxOpenConns)
	}

	if c.DatabaseMaxIdleConns < 0 || c.DatabaseMaxIdleConns > c.DatabaseMaxOpenConns {
		return fmt.Errorf(
			"DATABASE_MAX_IDLE_CONNS must be between 0 and DATABASE_MAX_OPEN_CONNS (%d): got %d",
			c.DatabaseMaxOpenConns, c.DatabaseMaxIdleConns,
		)
	}

	if err := c.validateIDMaskSecret(); err != nil {
		return err
	}

	return nil
}

func (c Config) validateIDMaskSecret() error {
	if c.IsLocal() {
		return nil
	}

	if len(c.IDMaskSecret) < minIDMaskSecretLength {
		return fmt.Errorf(
			"ID_MASK_SECRET must be at least %d characters outside local: got %d",
			minIDMaskSecretLength, len(c.IDMaskSecret),
		)
	}

	return nil
}

func env(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}

	return fallback
}

func envInt(key string, fallback int) (int, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback, nil
	}

	value, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer: got %q", key, raw)
	}

	return value, nil
}

func envInt64(key string, fallback int64) (int64, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback, nil
	}

	value, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer: got %q", key, raw)
	}

	return value, nil
}

func envUint64(key string, fallback uint64) (uint64, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback, nil
	}

	value, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be a non-negative integer: got %q", key, raw)
	}

	return value, nil
}

func envBool(key string, fallback bool) (bool, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback, nil
	}

	value, err := strconv.ParseBool(raw)
	if err != nil {
		return false, fmt.Errorf("%s must be true or false: got %q", key, raw)
	}

	return value, nil
}

func envDuration(key string, fallback time.Duration) (time.Duration, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback, nil
	}

	value, err := time.ParseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("%s must be a duration such as 10s or 1m: got %q", key, raw)
	}

	if value <= 0 {
		return 0, fmt.Errorf("%s must be greater than zero: got %q", key, raw)
	}

	return value, nil
}
