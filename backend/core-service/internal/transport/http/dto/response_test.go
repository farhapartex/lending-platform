package dto_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
)

func TestNewErrorResponse(t *testing.T) {
	response := dto.NewErrorResponse(dto.CodeNotFound, "The market does not exist.", "req-123")

	if response.Error.Code != dto.CodeNotFound {
		t.Fatalf("expected the code to be preserved, got %q", response.Error.Code)
	}

	if response.Error.Message != "The market does not exist." {
		t.Fatalf("unexpected message %q", response.Error.Message)
	}

	if response.RequestID != "req-123" {
		t.Fatalf("expected the request id, got %q", response.RequestID)
	}

	if response.Error.Details != nil {
		t.Fatalf("expected no details by default, got %v", response.Error.Details)
	}
}

func TestErrorResponseOmitsEmptyFields(t *testing.T) {
	encoded, err := json.Marshal(dto.NewErrorResponse(dto.CodeBadRequest, "Bad input.", ""))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `{"error":{"code":"BAD_REQUEST","message":"Bad input."}}`

	if string(encoded) != expected {
		t.Fatalf("expected %s, got %s", expected, encoded)
	}
}

func TestErrorCodesAreStable(t *testing.T) {
	cases := map[string]string{
		dto.CodeInternalError: "INTERNAL_ERROR",
		dto.CodeNotFound:      "NOT_FOUND",
		dto.CodeBadRequest:    "BAD_REQUEST",
	}

	for actual, expected := range cases {
		if actual != expected {
			t.Fatalf("expected the code %q to stay stable, got %q", expected, actual)
		}
	}
}

func TestNewHealthResponse(t *testing.T) {
	checkedAt := time.Date(2026, time.August, 18, 3, 25, 46, 0, time.UTC)

	response := dto.NewHealthResponse(domain.Health{
		Status:      domain.HealthStatusOK,
		ServiceName: "core-service",
		Version:     "v1.2.3",
		Environment: "local",
		Uptime:      90 * time.Second,
		CheckedAt:   checkedAt,
	})

	if response.Status != "ok" {
		t.Fatalf("expected ok, got %q", response.Status)
	}

	if response.Service != "core-service" || response.Version != "v1.2.3" || response.Environment != "local" {
		t.Fatalf("unexpected identity fields %+v", response)
	}

	if response.UptimeSeconds != 90 {
		t.Fatalf("expected 90 seconds, got %d", response.UptimeSeconds)
	}

	if response.CheckedAt != "2026-08-18T03:25:46Z" {
		t.Fatalf("expected an RFC 3339 timestamp, got %q", response.CheckedAt)
	}
}

func TestNewHealthResponseTruncatesPartialSeconds(t *testing.T) {
	response := dto.NewHealthResponse(domain.Health{
		Status:    domain.HealthStatusDegraded,
		Uptime:    1500 * time.Millisecond,
		CheckedAt: time.Now().UTC(),
	})

	if response.UptimeSeconds != 1 {
		t.Fatalf("expected whole seconds, got %d", response.UptimeSeconds)
	}

	if response.Status != "degraded" {
		t.Fatalf("expected degraded, got %q", response.Status)
	}
}

func TestNewHealthResponseConvertsToUTC(t *testing.T) {
	zone := time.FixedZone("UTC+6", 6*60*60)
	local := time.Date(2026, time.August, 18, 9, 25, 46, 0, zone)

	response := dto.NewHealthResponse(domain.Health{CheckedAt: local})

	if response.CheckedAt != "2026-08-18T03:25:46Z" {
		t.Fatalf("expected the timestamp converted to UTC, got %q", response.CheckedAt)
	}
}
