package transfer

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func CreateTestUser(ctx context.Context, sqlDB *sqlx.DB) (int64, error) {
	var id int64
	query, args, err := psql.Insert("users").
		Columns("username", "user_password").
		Values("test_user", "12345678").
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	q, a, err := psql.Insert("user_progress").
		Columns("user_id").
		Values(&id).
		ToSql()

	if err != nil {
		return 0, err
	}

	tx := sqlDB.MustBegin()
	tx.QueryRowxContext(ctx, query, args...).Scan(&id)
	tx.QueryxContext(ctx, q, a...)
	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}
