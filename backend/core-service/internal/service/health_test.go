package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/service"
)

func TestHealthCheckReportsIdentityAndUptime(t *testing.T) {
	startedAt := time.Date(2026, 8, 20, 3, 0, 0, 0, time.UTC)
	checkedAt := startedAt.Add(90 * time.Second)

	health := service.NewHealthService(service.HealthServiceParams{
		ServiceName: "core-service",
		Version:     "v1.2.3",
		Environment: "local",
		StartedAt:   startedAt,
		Now:         func() time.Time { return checkedAt },
	})

	result, err := health.Check(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Status != domain.HealthStatusOK {
		t.Fatalf("expected ok, got %q", result.Status)
	}

	if result.ServiceName != "core-service" || result.Version != "v1.2.3" || result.Environment != "local" {
		t.Fatalf("unexpected identity %+v", result)
	}

	if result.Uptime != 90*time.Second {
		t.Fatalf("expected 90s of uptime, got %s", result.Uptime)
	}

	if !result.CheckedAt.Equal(checkedAt) {
		t.Fatalf("expected the injected clock to be used, got %s", result.CheckedAt)
	}
}

func TestHealthUptimeGrowsWithTheClock(t *testing.T) {
	startedAt := time.Date(2026, 8, 20, 3, 0, 0, 0, time.UTC)
	current := startedAt

	health := service.NewHealthService(service.HealthServiceParams{
		StartedAt: startedAt,
		Now:       func() time.Time { return current },
	})

	first, err := health.Check(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	current = startedAt.Add(time.Hour)

	second, err := health.Check(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if second.Uptime <= first.Uptime {
		t.Fatalf("expected uptime to grow, got %s then %s", first.Uptime, second.Uptime)
	}
}

func TestHealthDefaultsTheClockWhenNoneIsSupplied(t *testing.T) {
	health := service.NewHealthService(service.HealthServiceParams{ServiceName: "core-service"})

	result, err := health.Check(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.CheckedAt.IsZero() {
		t.Fatal("expected a real timestamp from the default clock")
	}

	if result.Uptime < 0 {
		t.Fatalf("expected a non negative uptime, got %s", result.Uptime)
	}
}

func TestHealthDefaultsStartedAtToTheFirstClockReading(t *testing.T) {
	moment := time.Date(2026, 8, 20, 3, 0, 0, 0, time.UTC)

	health := service.NewHealthService(service.HealthServiceParams{
		Now: func() time.Time { return moment },
	})

	result, err := health.Check(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Uptime != 0 {
		t.Fatalf("expected zero uptime when started at the current instant, got %s", result.Uptime)
	}
}

func TestHealthCheckRespectsACancelledContext(t *testing.T) {
	health := service.NewHealthService(service.HealthServiceParams{ServiceName: "core-service"})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	result, err := health.Check(ctx)
	if err == nil {
		t.Fatal("expected a cancelled context to be reported")
	}

	if result.Status != "" {
		t.Fatalf("expected an empty result on error, got %+v", result)
	}
}
