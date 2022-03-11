package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
)

func NewSqlite(dsn, driver string) (*sqlx.DB, error) {
	ctx := context.Background()
	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		logger.ErrorKV(ctx, "NewSqlite -- failed to create database connection", "error", err)

		return nil, err
	}

	return db, nil
}
