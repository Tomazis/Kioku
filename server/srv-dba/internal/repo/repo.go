package repo

import (
	"sync"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type repo struct {
	db    *sqlx.DB
	mutex sync.Mutex
}

func NewRepo(db_ *sqlx.DB) *repo {
	return &repo{db: db_}
}
