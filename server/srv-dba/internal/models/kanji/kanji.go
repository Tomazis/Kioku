package models

import "database/sql"

type Kanji struct {
	ID           uint64         `db:"id"`
	Kanji        string         `db:"kanji"`
	Primary      string         `db:"kanji_meaning"`
	Level        uint32         `db:"kanji_level"`
	Alternatives sql.NullString `db:"kanji_alternative"`
	Onyomi       sql.NullString `db:"kanji_onyomi"`
	Kunyomi      sql.NullString `db:"kanji_kunyomi"`
}

type KanjiAlternative struct {
	ID          uint64 `db:"id"`
	KanjiID     uint64 `db:"kanji_id"`
	Alternative string `db:"kanji_alternative"`
}

type Onyomi struct {
	ID      uint64 `db:"id"`
	KanjiID uint64 `db:"kanji_id"`
	Onyomi  string `db:"kanji_onyomi"`
}

type Kunyomi struct {
	ID      uint64 `db:"id"`
	KanjiID uint64 `db:"kanji_id"`
	Kunyomi string `db:"kanji_kunyomi"`
}
