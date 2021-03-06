package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/tomazis/kioku/server/srv-dba/internal/logger"

	_ "github.com/lib/pq"
)

func NewPostgres(dsn, driver string) (*sqlx.DB, error) {
	ctx := context.Background()
	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		logger.ErrorKV(ctx, "NewPostgres -- failed to create database connection", "error", err)

		return nil, err
	}

	return db, nil
}
