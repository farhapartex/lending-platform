package http

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/config"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/handler"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/middleware"
)

const APIBasePath = "/api/v1"

type RouterParams struct {
	Config        config.Config
	Logger        *slog.Logger
	HealthService domain.HealthService
}

func NewRouter(params RouterParams) *gin.Engine {
	if params.Config.IsLocal() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.RedirectTrailingSlash = false
	engine.HandleMethodNotAllowed = true

	engine.Use(
		middleware.RequestID(),
		middleware.Logger(params.Logger),
		middleware.Recovery(params.Logger),
	)

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(
			http.StatusNotFound,
			dto.NewErrorResponse(dto.CodeNotFound, "The requested endpoint does not exist.", middleware.RequestIDFrom(c)),
		)
	})

	engine.NoMethod(func(c *gin.Context) {
		c.JSON(
			http.StatusMethodNotAllowed,
			dto.NewErrorResponse(dto.CodeBadRequest, "That method is not allowed on this endpoint.", middleware.RequestIDFrom(c)),
		)
	})

	healthHandler := handler.NewHealthHandler(params.HealthService)

	v1 := engine.Group(APIBasePath)
	v1.GET("/health", healthHandler.Get)

	return engine
}
