package models

import models "github.com/tomazis/kioku/server/srv-frontend-api/internal/models/kanji"

type Word struct {
	ID           uint64            `db:"id"`
	Word         string            `db:"word"`
	Primary      string            `db:"primary"`
	Level        uint32            `db:"level"`
	Composition  []models.Kanji    `db:"-"`
	Alternatives []WordAlternative `db:"-"`
	Readings     []WordReading     `db:"-"`
	Types        []WordType        `db:"-"`
	Sentences    []Sentence        `db:"-"`
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
