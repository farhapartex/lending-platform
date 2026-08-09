package domain

import (
	"context"
	"time"
)

type HealthStatus string

const (
	HealthStatusOK       HealthStatus = "ok"
	HealthStatusDegraded HealthStatus = "degraded"
	HealthStatusDown     HealthStatus = "down"
)

type Health struct {
	Status      HealthStatus
	ServiceName string
	Version     string
	Environment string
	Uptime      time.Duration
	CheckedAt   time.Time
}

type HealthService interface {
	Check(ctx context.Context) (Health, error)
}
