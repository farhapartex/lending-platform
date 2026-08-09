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
)

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
	}

	port, err := envInt("HTTP_PORT", 8080)
	if err != nil {
		return Config{}, err
	}
	cfg.HTTPPort = port

	durations := []struct {
		key    string
		fallb  time.Duration
		target *time.Duration
	}{
		{"HTTP_READ_TIMEOUT", 10 * time.Second, &cfg.ReadTimeout},
		{"HTTP_WRITE_TIMEOUT", 15 * time.Second, &cfg.WriteTimeout},
		{"HTTP_IDLE_TIMEOUT", 60 * time.Second, &cfg.IdleTimeout},
		{"SHUTDOWN_GRACE_PERIOD", 15 * time.Second, &cfg.ShutdownGrace},
	}

	for _, d := range durations {
		value, err := envDuration(d.key, d.fallb)
		if err != nil {
			return Config{}, err
		}
		*d.target = value
	}

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
