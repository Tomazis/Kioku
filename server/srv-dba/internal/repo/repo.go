package repo

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	db    *sqlx.DB
	mutex sync.Mutex
}

func NewRepo(db_ *sqlx.DB) *repo {
	return &repo{db: db_}
}
