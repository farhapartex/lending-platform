package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/middleware"
)

type HealthHandler struct {
	healthService domain.HealthService
}

func NewHealthHandler(healthService domain.HealthService) *HealthHandler {
	return &HealthHandler{healthService: healthService}
}

func (h *HealthHandler) Get(c *gin.Context) {
	health, err := h.healthService.Check(c.Request.Context())
	if err != nil {
		c.JSON(
			http.StatusServiceUnavailable,
			dto.NewErrorResponse(dto.CodeInternalError, "Health check could not be completed.", middleware.RequestIDFrom(c)),
		)

		return
	}

	c.JSON(http.StatusOK, dto.NewHealthResponse(health))
}
