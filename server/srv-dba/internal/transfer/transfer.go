package transfer

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
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
	_, err := CreateTestUser(ctx, pgDB)

	if err != nil {
		return err
	}

	return nil

	_, err = GetSqliteKanji(ctx, sqliteDB)

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
