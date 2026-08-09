package logger

import (
	"log/slog"
	"os"
	"strings"
)

func New(level, environment, serviceName, version string) *slog.Logger {
	options := &slog.HandlerOptions{Level: parseLevel(level)}

	var handler slog.Handler
	if environment == "local" {
		handler = slog.NewTextHandler(os.Stdout, options)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, options)
	}

	return slog.New(handler).With(
		slog.String("service", serviceName),
		slog.String("version", version),
		slog.String("environment", environment),
	)
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
