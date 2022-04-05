package models

import (
	"database/sql"

	m_kanji "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
)

type Word struct {
	ID                   uint64           `db:"id"`
	Word                 string           `db:"word"`
	Primary              string           `db:"word_meaning"`
	Level                uint32           `db:"word_level"`
	Composition          []*m_kanji.Kanji `db:"-"`
	Alternatives         sql.NullString   `db:"word_alternative"`
	Readings             sql.NullString   `db:"word_reading"`
	Types                sql.NullString   `db:"word_type"`
	Sentences            sql.NullString   `db:"japanese_sentence"`
	SentenceTranslations sql.NullString   `db:"sentence_translation"`
	SentenceLanguage     sql.NullString   `db:"sentence_language"`
}

type WordAlternative struct {
	ID          uint64 `db:"id"`
	WordID      uint64 `db:"word_id"`
	Alternative string `db:"alternative"`
}

type WordReading struct {
	ID      uint64 `db:"id"`
	WordID  uint64 `db:"word_id"`
	Reading string `db:"reading"`
}

type WordType struct {
	ID     uint64 `db:"id"`
	WordID uint64 `db:"word_id"`
	Type   string `db:"type"`
}
