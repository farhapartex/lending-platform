package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/handler"
)

type stubHealthService struct {
	health   domain.Health
	failWith error
}

func (s *stubHealthService) Check(context.Context) (domain.Health, error) {
	if s.failWith != nil {
		return domain.Health{}, s.failWith
	}

	return s.health, nil
}

func newHealthRouter(service domain.HealthService) *gin.Engine {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	engine.GET("/health", handler.NewHealthHandler(service).Get)

	return engine
}

func TestHealthHandlerReturnsTheReport(t *testing.T) {
	engine := newHealthRouter(&stubHealthService{health: domain.Health{
		Status:      domain.HealthStatusOK,
		ServiceName: "core-service",
		Version:     "v1.2.3",
		Environment: "local",
		Uptime:      90 * time.Second,
		CheckedAt:   time.Date(2026, 8, 20, 3, 25, 46, 0, time.UTC),
	}})

	recorder := performGet(engine, "/health")

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	var response dto.HealthResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("could not decode the body: %v", err)
	}

	if response.Status != "ok" || response.Service != "core-service" {
		t.Fatalf("unexpected payload %+v", response)
	}

	if response.UptimeSeconds != 90 {
		t.Fatalf("expected 90 seconds, got %d", response.UptimeSeconds)
	}

	if response.CheckedAt != "2026-08-20T03:25:46Z" {
		t.Fatalf("unexpected timestamp %q", response.CheckedAt)
	}
}

func TestHealthHandlerReportsFailureAsUnavailable(t *testing.T) {
	engine := newHealthRouter(&stubHealthService{failWith: errors.New("clock unavailable")})

	recorder := performGet(engine, "/health")

	if recorder.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503, got %d", recorder.Code)
	}

	body := decodeError(t, recorder.Body.Bytes())

	if body.Error.Code != dto.CodeInternalError {
		t.Fatalf("expected INTERNAL_ERROR, got %s", recorder.Body)
	}

	if contains(body.Error.Message, "clock") {
		t.Fatalf("expected the internal detail to stay hidden, got %q", body.Error.Message)
	}
}
