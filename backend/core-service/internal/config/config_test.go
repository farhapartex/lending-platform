package config

import (
	"strings"
	"testing"
	"time"
)

func TestIsLocal(t *testing.T) {
	cases := []struct {
		appEnv string
		want   bool
	}{
		{appEnv: EnvLocal, want: true},
		{appEnv: EnvStaging, want: false},
		{appEnv: EnvProduction, want: false},
		{appEnv: "", want: false},
	}

	for _, testCase := range cases {
		t.Run(testCase.appEnv, func(t *testing.T) {
			cfg := Config{AppEnv: testCase.appEnv}

			if got := cfg.IsLocal(); got != testCase.want {
				t.Fatalf("expected %v, got %v", testCase.want, got)
			}
		})
	}
}

func TestLoadDefaults(t *testing.T) {
	setRequiredChainEnv(t)

	cfg, err := Load("v1.2.3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.AppEnv != EnvLocal {
		t.Fatalf("expected the default environment to be local, got %q", cfg.AppEnv)
	}

	if cfg.ServiceName != "core-service" {
		t.Fatalf("expected the default service name, got %q", cfg.ServiceName)
	}

	if cfg.ServiceVersion != "v1.2.3" {
		t.Fatalf("expected the version to come from the argument, got %q", cfg.ServiceVersion)
	}

	if cfg.HTTPPort != 8080 {
		t.Fatalf("expected the default port 8080, got %d", cfg.HTTPPort)
	}

	if cfg.LogLevel != "info" {
		t.Fatalf("expected the default log level info, got %q", cfg.LogLevel)
	}

	if cfg.DatabaseMaxOpenConns != 20 || cfg.DatabaseMaxIdleConns != 10 {
		t.Fatalf("expected default pool sizes, got open=%d idle=%d", cfg.DatabaseMaxOpenConns, cfg.DatabaseMaxIdleConns)
	}

	if cfg.DatabaseLogQueries {
		t.Fatal("expected query logging to default to off")
	}

	if cfg.ReadTimeout != 10*time.Second || cfg.WriteTimeout != 15*time.Second {
		t.Fatalf("expected default http timeouts, got read=%s write=%s", cfg.ReadTimeout, cfg.WriteTimeout)
	}

	if cfg.IdleTimeout != 60*time.Second || cfg.ShutdownGrace != 15*time.Second {
		t.Fatalf("expected default idle and shutdown values, got idle=%s grace=%s", cfg.IdleTimeout, cfg.ShutdownGrace)
	}

	if cfg.DatabaseConnMaxLifetime != 30*time.Minute || cfg.DatabaseConnMaxIdleTime != 5*time.Minute {
		t.Fatalf("expected default connection lifetimes, got %s and %s", cfg.DatabaseConnMaxLifetime, cfg.DatabaseConnMaxIdleTime)
	}
}

func TestLoadReadsOverrides(t *testing.T) {
	setRequiredChainEnv(t)

	t.Setenv("APP_ENV", EnvProduction)
	t.Setenv("SERVICE_NAME", "lending-api")
	t.Setenv("LOG_LEVEL", "warn")
	t.Setenv("HTTP_PORT", "9090")
	t.Setenv("DATABASE_URL", "postgres://localhost:5432/lending")
	t.Setenv("DATABASE_MAX_OPEN_CONNS", "40")
	t.Setenv("DATABASE_MAX_IDLE_CONNS", "5")
	t.Setenv("DATABASE_LOG_QUERIES", "true")
	t.Setenv("HTTP_READ_TIMEOUT", "30s")
	t.Setenv("ID_MASK_SECRET", strings.Repeat("s", minIDMaskSecretLength))

	cfg, err := Load("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.AppEnv != EnvProduction || cfg.ServiceName != "lending-api" || cfg.LogLevel != "warn" {
		t.Fatalf("expected overrides to apply, got %+v", cfg)
	}

	if cfg.HTTPPort != 9090 {
		t.Fatalf("expected port 9090, got %d", cfg.HTTPPort)
	}

	if cfg.DatabaseURL != "postgres://localhost:5432/lending" {
		t.Fatalf("expected the database url to apply, got %q", cfg.DatabaseURL)
	}

	if !cfg.DatabaseLogQueries {
		t.Fatal("expected query logging to be enabled")
	}

	if cfg.ReadTimeout != 30*time.Second {
		t.Fatalf("expected a 30s read timeout, got %s", cfg.ReadTimeout)
	}
}

func TestLoadRejectsBadEnvValues(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		value    string
		contains string
	}{
		{name: "port not a number", key: "HTTP_PORT", value: "http", contains: "HTTP_PORT"},
		{name: "port too low", key: "HTTP_PORT", value: "0", contains: "HTTP_PORT"},
		{name: "port too high", key: "HTTP_PORT", value: "70000", contains: "HTTP_PORT"},
		{name: "max open not a number", key: "DATABASE_MAX_OPEN_CONNS", value: "many", contains: "DATABASE_MAX_OPEN_CONNS"},
		{name: "max open zero", key: "DATABASE_MAX_OPEN_CONNS", value: "0", contains: "DATABASE_MAX_OPEN_CONNS"},
		{name: "max idle not a number", key: "DATABASE_MAX_IDLE_CONNS", value: "some", contains: "DATABASE_MAX_IDLE_CONNS"},
		{name: "max idle negative", key: "DATABASE_MAX_IDLE_CONNS", value: "-1", contains: "DATABASE_MAX_IDLE_CONNS"},
		{name: "max idle above max open", key: "DATABASE_MAX_IDLE_CONNS", value: "999", contains: "DATABASE_MAX_IDLE_CONNS"},
		{name: "log queries not a bool", key: "DATABASE_LOG_QUERIES", value: "sometimes", contains: "DATABASE_LOG_QUERIES"},
		{name: "read timeout not a duration", key: "HTTP_READ_TIMEOUT", value: "quick", contains: "HTTP_READ_TIMEOUT"},
		{name: "read timeout zero", key: "HTTP_READ_TIMEOUT", value: "0s", contains: "HTTP_READ_TIMEOUT"},
		{name: "shutdown grace negative", key: "SHUTDOWN_GRACE_PERIOD", value: "-1s", contains: "SHUTDOWN_GRACE_PERIOD"},
		{name: "unknown environment", key: "APP_ENV", value: "qa", contains: "APP_ENV"},
		{name: "unknown log level", key: "LOG_LEVEL", value: "verbose", contains: "LOG_LEVEL"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			setRequiredChainEnv(t)
			t.Setenv(testCase.key, testCase.value)

			_, err := Load("test")
			if err == nil {
				t.Fatalf("expected an error for %s=%q", testCase.key, testCase.value)
			}

			if !strings.Contains(err.Error(), testCase.contains) {
				t.Fatalf("expected the error to mention %s, got %v", testCase.contains, err)
			}
		})
	}
}

func TestLoadPropagatesChainErrors(t *testing.T) {
	setRequiredChainEnv(t)
	t.Setenv("CHAIN_ID", "not-a-chain")

	_, err := Load("test")
	if err == nil {
		t.Fatal("expected a chain configuration error to surface from Load")
	}

	if !strings.Contains(err.Error(), "CHAIN_ID") {
		t.Fatalf("expected the error to mention CHAIN_ID, got %v", err)
	}
}

func TestLogLevelIsCaseInsensitive(t *testing.T) {
	setRequiredChainEnv(t)
	t.Setenv("LOG_LEVEL", "DEBUG")

	cfg, err := Load("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.LogLevel != "DEBUG" {
		t.Fatalf("expected the raw value to be preserved, got %q", cfg.LogLevel)
	}
}

func TestValidateRejectsEmptyServiceName(t *testing.T) {
	cfg := Config{
		AppEnv:               EnvLocal,
		LogLevel:             "info",
		HTTPPort:             8080,
		ServiceName:          "",
		DatabaseMaxOpenConns: 10,
		DatabaseMaxIdleConns: 5,
	}

	err := cfg.validate()
	if err == nil {
		t.Fatal("expected an error for an empty service name")
	}

	if !strings.Contains(err.Error(), "SERVICE_NAME") {
		t.Fatalf("expected the error to mention SERVICE_NAME, got %v", err)
	}
}

func TestIDMaskSecretIsOptionalLocally(t *testing.T) {
	setRequiredChainEnv(t)

	cfg, err := Load("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.IDMaskSecret != "" {
		t.Fatalf("expected no configured secret locally, got %q", cfg.IDMaskSecret)
	}

	if cfg.EffectiveIDMaskSecret() != localIDMaskSecret {
		t.Fatalf("expected the local fallback secret, got %q", cfg.EffectiveIDMaskSecret())
	}
}

func TestIDMaskSecretIsRequiredOutsideLocal(t *testing.T) {
	for _, appEnv := range []string{EnvStaging, EnvProduction} {
		t.Run(appEnv, func(t *testing.T) {
			setRequiredChainEnv(t)
			t.Setenv("APP_ENV", appEnv)

			_, err := Load("test")
			if err == nil {
				t.Fatalf("expected an error when the secret is missing in %s", appEnv)
			}

			if !strings.Contains(err.Error(), "ID_MASK_SECRET") {
				t.Fatalf("expected the error to mention ID_MASK_SECRET, got %v", err)
			}
		})
	}
}

func TestIDMaskSecretRejectsShortValuesOutsideLocal(t *testing.T) {
	setRequiredChainEnv(t)
	t.Setenv("APP_ENV", EnvProduction)
	t.Setenv("ID_MASK_SECRET", strings.Repeat("a", minIDMaskSecretLength-1))

	_, err := Load("test")
	if err == nil {
		t.Fatal("expected a short secret to be rejected")
	}

	if !strings.Contains(err.Error(), "ID_MASK_SECRET") {
		t.Fatalf("expected the error to mention ID_MASK_SECRET, got %v", err)
	}
}

func TestIDMaskSecretAcceptedWhenLongEnough(t *testing.T) {
	setRequiredChainEnv(t)
	t.Setenv("APP_ENV", EnvProduction)

	secret := strings.Repeat("s", minIDMaskSecretLength)
	t.Setenv("ID_MASK_SECRET", secret)

	cfg, err := Load("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.EffectiveIDMaskSecret() != secret {
		t.Fatalf("expected the configured secret to be used, got %q", cfg.EffectiveIDMaskSecret())
	}
}

func TestCacheSettingsDefaults(t *testing.T) {
	setRequiredChainEnv(t)

	cfg, err := Load("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.CacheEnabled() {
		t.Fatal("expected the cache to be disabled when REDIS_URL is unset")
	}

	if cfg.CacheNamespace != cfg.ServiceName {
		t.Fatalf("expected the namespace to default to the service name, got %q", cfg.CacheNamespace)
	}

	if cfg.CacheRequestTimeout != 2*time.Second {
		t.Fatalf("expected a 2s default cache timeout, got %s", cfg.CacheRequestTimeout)
	}
}

func TestCacheSettingsOverrides(t *testing.T) {
	setRequiredChainEnv(t)

	t.Setenv("REDIS_URL", "redis://localhost:6379/0")
	t.Setenv("CACHE_NAMESPACE", "lending-cache")
	t.Setenv("CACHE_REQUEST_TIMEOUT", "5s")

	cfg, err := Load("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !cfg.CacheEnabled() {
		t.Fatal("expected the cache to be enabled")
	}

	if cfg.RedisURL != "redis://localhost:6379/0" {
		t.Fatalf("unexpected redis url %q", cfg.RedisURL)
	}

	if cfg.CacheNamespace != "lending-cache" {
		t.Fatalf("unexpected namespace %q", cfg.CacheNamespace)
	}

	if cfg.CacheRequestTimeout != 5*time.Second {
		t.Fatalf("unexpected cache timeout %s", cfg.CacheRequestTimeout)
	}
}

func TestCacheRequestTimeoutMustBeValid(t *testing.T) {
	setRequiredChainEnv(t)
	t.Setenv("CACHE_REQUEST_TIMEOUT", "instantly")

	_, err := Load("test")
	if err == nil {
		t.Fatal("expected an error for an invalid cache timeout")
	}

	if !strings.Contains(err.Error(), "CACHE_REQUEST_TIMEOUT") {
		t.Fatalf("expected the error to mention CACHE_REQUEST_TIMEOUT, got %v", err)
	}
}

func TestEnvHelpersFallBackOnBlankValues(t *testing.T) {
	t.Setenv("BLANK_STRING", "   ")
	if got := env("BLANK_STRING", "fallback"); got != "fallback" {
		t.Fatalf("expected the fallback, got %q", got)
	}

	t.Setenv("BLANK_INT", "  ")
	value, err := envInt("BLANK_INT", 7)
	if err != nil || value != 7 {
		t.Fatalf("expected fallback 7, got %d with error %v", value, err)
	}

	t.Setenv("BLANK_INT64", "  ")
	value64, err := envInt64("BLANK_INT64", 8)
	if err != nil || value64 != 8 {
		t.Fatalf("expected fallback 8, got %d with error %v", value64, err)
	}

	t.Setenv("BLANK_UINT64", "  ")
	valueUint, err := envUint64("BLANK_UINT64", 9)
	if err != nil || valueUint != 9 {
		t.Fatalf("expected fallback 9, got %d with error %v", valueUint, err)
	}

	t.Setenv("BLANK_BOOL", "  ")
	valueBool, err := envBool("BLANK_BOOL", true)
	if err != nil || !valueBool {
		t.Fatalf("expected fallback true, got %v with error %v", valueBool, err)
	}

	t.Setenv("BLANK_DURATION", "  ")
	valueDuration, err := envDuration("BLANK_DURATION", time.Minute)
	if err != nil || valueDuration != time.Minute {
		t.Fatalf("expected fallback 1m, got %s with error %v", valueDuration, err)
	}
}
