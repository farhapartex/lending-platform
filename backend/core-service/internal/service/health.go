package service

import (
	"context"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

type HealthServiceParams struct {
	ServiceName string
	Version     string
	Environment string
	StartedAt   time.Time
	Now         func() time.Time
}

type healthService struct {
	serviceName string
	version     string
	environment string
	startedAt   time.Time
	now         func() time.Time
}

func NewHealthService(params HealthServiceParams) domain.HealthService {
	now := params.Now
	if now == nil {
		now = time.Now
	}

	startedAt := params.StartedAt
	if startedAt.IsZero() {
		startedAt = now()
	}

	return &healthService{
		serviceName: params.ServiceName,
		version:     params.Version,
		environment: params.Environment,
		startedAt:   startedAt,
		now:         now,
	}
}

func (s *healthService) Check(ctx context.Context) (domain.Health, error) {
	if err := ctx.Err(); err != nil {
		return domain.Health{}, err
	}

	checkedAt := s.now()

	return domain.Health{
		Status:      domain.HealthStatusOK,
		ServiceName: s.serviceName,
		Version:     s.version,
		Environment: s.environment,
		Uptime:      checkedAt.Sub(s.startedAt),
		CheckedAt:   checkedAt,
	}, nil
}
