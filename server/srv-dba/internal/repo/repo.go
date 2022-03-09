package repo

import "github.com/jmoiron/sqlx"

type repo struct {
	db *sqlx.DB
}

func NewRepo(db_ *sqlx.DB) *repo {
	return &repo{db: db_}
}
