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

type wordModel struct {
	ID              uint64         `db:"id"`
	Word            string         `db:"word"`
	WordMeaning     string         `db:"word_meaning"`
	WordLevel       uint32         `db:"word_level"`
	WordReading     string         `db:"word_reading"`
	WordType        string         `db:"word_type"`
	SentenceJap     string         `db:"jap"`
	SentenceEng     string         `db:"eng"`
	WordAlternative sql.NullString `db:"word_alternative"`
	Progress        sql.NullString `db:"progress"`
}

type compostion struct {
	WordID  uint64 `db:"word_id"`
	KanjiID uint64 `db:"kanji_id"`
}

func Transfer(ctx context.Context, sqliteDB *sqlx.DB, pgDB *sqlx.DB) error {
	_, err := GetSqliteKanji(ctx, sqliteDB)

	if err != nil {
		return err
	}

	_, err = GetSqliteWords(ctx, sqliteDB)

	if err != nil {
		return err
	}

	_, err = GetSqliteCompositions(ctx, sqliteDB)

	if err != nil {
		return err
	}

	return nil
}

func GetSqliteKanji(ctx context.Context, sqlDB *sqlx.DB) ([]kanjiModel, error) {
	query, _, err := sq.Select("kanji.*, kanji_alternatives.kanji_alternative, kanji_kunyomi.kanji_kunyomi, kanji_onyomi.kanji_onyomi").
		From("kanji").
		LeftJoin("kanji_alternatives ON (kanji.id == kanji_alternatives.kanji_id)").
		LeftJoin("kanji_kunyomi ON (kanji.id == kanji_kunyomi.kanji_id)").
		LeftJoin("kanji_onyomi ON (kanji.id == kanji_onyomi.kanji_id)").
		ToSql()

	if err != nil {
		return nil, err
	}

	logger.InfoKV(ctx, "Kanji query", "query", query)

	var kanji []kanjiModel

	err = sqlDB.SelectContext(ctx, &kanji, query)

	if err != nil {
		return nil, err
	}

	return kanji, err
}

func GetSqliteWords(ctx context.Context, sqlDB *sqlx.DB) ([]wordModel, error) {
	query, _, err := sq.Select("words.*, word_alternatives.word_alternative, word_readings.word_reading, word_types.word_type, sentences.jap, sentences.eng").
		From("words").
		LeftJoin("word_alternatives ON (words.id == word_alternatives.word_id)").
		LeftJoin("word_readings ON (words.id == word_readings.word_id)").
		LeftJoin("word_types ON (words.id == word_types.word_id)").
		LeftJoin("sentences ON (words.id == sentences.word_id)").
		ToSql()

	if err != nil {
		return nil, err
	}

	logger.InfoKV(ctx, "Words query", "query", query)

	var words []wordModel

	err = sqlDB.SelectContext(ctx, &words, query)

	if err != nil {
		return nil, err
	}

	logger.InfoKV(ctx, "Words count", "count", len(words))

	return words, err
}

func GetSqliteCompositions(ctx context.Context, sqlDB *sqlx.DB) ([]compostion, error) {
	query, _, err := sq.Select("compositions.*").
		From("compositions").ToSql()

	if err != nil {
		return nil, err
	}

	logger.InfoKV(ctx, "Compostions query", "query", query)

	var compostions []compostion

	err = sqlDB.SelectContext(ctx, &compostions, query)

	if err != nil {
		return nil, err
	}

	logger.InfoKV(ctx, "Words count", "count", len(compostions))

	return compostions, err
}
