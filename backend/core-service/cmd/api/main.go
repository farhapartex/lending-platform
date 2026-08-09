package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/config"
	"github.com/farhapartex/lending-platform/core-service/internal/platform/logger"
	"github.com/farhapartex/lending-platform/core-service/internal/service"
	transporthttp "github.com/farhapartex/lending-platform/core-service/internal/transport/http"
)

var (
	version = "dev"
	commit  = "unknown"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "core-service failed to start: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	startedAt := time.Now()

	cfg, err := config.Load(version)
	if err != nil {
		return err
	}

	log := logger.New(cfg.LogLevel, cfg.AppEnv, cfg.ServiceName, cfg.ServiceVersion)

	healthService := service.NewHealthService(service.HealthServiceParams{
		ServiceName: cfg.ServiceName,
		Version:     cfg.ServiceVersion,
		Environment: cfg.AppEnv,
		StartedAt:   startedAt,
	})

	router := transporthttp.NewRouter(transporthttp.RouterParams{
		Config:        cfg,
		Logger:        log,
		HealthService: healthService,
	})

	server := &http.Server{
		Addr:              net.JoinHostPort("", strconv.Itoa(cfg.HTTPPort)),
		Handler:           router,
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serverErrors := make(chan error, 1)

	go func() {
		log.Info(
			"http server listening",
			slog.Int("port", cfg.HTTPPort),
			slog.String("commit", commit),
			slog.String("health_endpoint", transporthttp.APIBasePath+"/health"),
		)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("http server stopped unexpectedly: %w", err)
	case <-ctx.Done():
		log.Info("shutdown signal received", slog.Duration("grace_period", cfg.ShutdownGrace))
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownGrace)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}

	log.Info("shutdown complete")

	return nil
}
