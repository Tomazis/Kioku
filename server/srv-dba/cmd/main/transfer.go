package main

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"

	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
)

type kanjiModel struct {
	ID          uint64         `db:"id"`
	Kanji       string         `db:"kanji"`
	Primary     string         `db:"kanji_meaning"`
	Level       uint32         `db:"kanji_level"`
	Alternative sql.NullString `db:"kanji_alternative"`
	Kunyomi     sql.NullString `db:"kanji_kunyomi"`
	Onyomi      sql.NullString `db:"kanji_onyomi"`
	Progress    sql.NullString `db:"progress"`
}

func Transfer(ctx context.Context, sqliteDB *sqlx.DB, pgDB *sqlx.DB) error {
	_, err := GetSqliteKanji(ctx, sqliteDB)

	if err != nil {
		return err
	}

	return nil
}

func GetSqliteKanji(ctx context.Context, sqlDB *sqlx.DB) ([]*kanjiModel, error) {

	query, _, err := sq.Select("kanji.*, kanji_alternative.kanji_alternative, kanji_kunyomi.kanji_kunyomi, kanji_onyomi.kanji_onyomi").
		From("kanji").
		LeftJoin("kanji_alternative ON (kanji.id == kanji_alternative.kanji_id)").
		LeftJoin("kanji_kunyomi ON (kanji.id == kanji_kunyomi.kanji_id)").
		LeftJoin("kanji_onyomi ON (kanji.id == kanji_onyomi.kanji_id)").
		ToSql()

	logger.InfoKV(ctx, "Kanji query", "query", query)

	if err != nil {
		return nil, err
	}

	var kanji []kanjiModel

	err = sqlDB.SelectContext(ctx, &kanji, query)

	if err != nil {
		return nil, err
	}

	return nil, err
}
