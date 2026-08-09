package dto

import (
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

type HealthResponse struct {
	Status        string `json:"status"`
	Service       string `json:"service"`
	Version       string `json:"version"`
	Environment   string `json:"environment"`
	UptimeSeconds int64  `json:"uptime_seconds"`
	CheckedAt     string `json:"checked_at"`
}

func NewHealthResponse(health domain.Health) HealthResponse {
	return HealthResponse{
		Status:        string(health.Status),
		Service:       health.ServiceName,
		Version:       health.Version,
		Environment:   health.Environment,
		UptimeSeconds: int64(health.Uptime.Seconds()),
		CheckedAt:     health.CheckedAt.UTC().Format(time.RFC3339),
	}
}
