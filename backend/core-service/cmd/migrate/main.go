package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/farhapartex/lending-platform/core-service/internal/config"
	"github.com/farhapartex/lending-platform/core-service/internal/platform/database"
	"github.com/farhapartex/lending-platform/core-service/migrations"
)

const migrationTimeout = 5 * time.Minute

func main() {
	command := flag.String("command", "up", "up, down, redo, status, version, reset")
	flag.Parse()

	if err := run(*command); err != nil {
		fmt.Fprintf(os.Stderr, "migrate failed: %v\n", err)
		os.Exit(1)
	}
}

func run(command string) error {
	cfg, err := config.Load("migrate")
	if err != nil {
		return err
	}

	if cfg.DatabaseURL == "" {
		return errors.New("DATABASE_URL must be set to run migrations")
	}

	pool, err := database.OpenSQL(database.Options{
		DSN:          cfg.DatabaseURL,
		MaxOpenConns: 2,
		MaxIdleConns: 1,
	})
	if err != nil {
		return err
	}

	defer func() {
		_ = pool.Close()
	}()

	goose.SetBaseFS(migrations.FS)
	goose.SetTableName("schema_migrations")

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose dialect: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), migrationTimeout)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		return fmt.Errorf("database unreachable: %w", err)
	}

	switch command {
	case "up":
		return goose.UpContext(ctx, pool, ".")
	case "down":
		return goose.DownContext(ctx, pool, ".")
	case "redo":
		return goose.RedoContext(ctx, pool, ".")
	case "reset":
		return goose.ResetContext(ctx, pool, ".")
	case "status":
		return goose.StatusContext(ctx, pool, ".")
	case "version":
		return goose.VersionContext(ctx, pool, ".")
	default:
		return fmt.Errorf("unknown command %q", command)
	}
}
