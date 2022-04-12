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

func prepareCompStatement(whatSelect string, whereSq interface{}, args ...interface{}) (string, []interface{}, error) {
	q, args, err := psql.Select(whatSelect).
		From("compositions").
		Where(whereSq).ToSql()

	return q, args, err
}
