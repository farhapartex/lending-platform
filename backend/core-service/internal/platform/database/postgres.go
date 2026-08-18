package database

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Options struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	LogQueries      bool
}

func Open(opts Options) (*gorm.DB, error) {
	gormLogLevel := logger.Warn
	if opts.LogQueries {
		gormLogLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(opts.DSN), &gorm.Config{
		Logger:                 logger.Default.LogMode(gormLogLevel),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		DisableAutomaticPing: false,
		TranslateError:       true,
	})
	if err != nil {
		return nil, fmt.Errorf("database: open failed: %w", err)
	}

	pool, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("database: pool unavailable: %w", err)
	}

	applyPoolSettings(pool, opts)

	return db, nil
}

func OpenSQL(opts Options) (*sql.DB, error) {
	pool, err := sql.Open("pgx", opts.DSN)
	if err != nil {
		return nil, fmt.Errorf("database: sql open failed: %w", err)
	}

	applyPoolSettings(pool, opts)

	return pool, nil
}

func applyPoolSettings(pool *sql.DB, opts Options) {
	if opts.MaxOpenConns > 0 {
		pool.SetMaxOpenConns(opts.MaxOpenConns)
	}

	if opts.MaxIdleConns > 0 {
		pool.SetMaxIdleConns(opts.MaxIdleConns)
	}

	if opts.ConnMaxLifetime > 0 {
		pool.SetConnMaxLifetime(opts.ConnMaxLifetime)
	}

	if opts.ConnMaxIdleTime > 0 {
		pool.SetConnMaxIdleTime(opts.ConnMaxIdleTime)
	}
}
