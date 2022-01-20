package models

type Kanji struct {
	ID           uint64   `db:"id"`
	Kanji        string   `db:"kanji"`
	Primary      string   `db:"primary"`
	Level        uint32   `db:"level"`
	Alternatives []string `db:"-"`
	Onyomi       []string `db:"-"`
	Kunyomi      []string `db:"-"`
}

type KanjiAlternative struct {
	ID          uint64 `db:"id"`
	KanjiID     uint64 `db:"kanji_id"`
	Alternative string `db:"alternative"`
}

type Onyomi struct {
	ID      uint64 `db:"id"`
	KanjiID uint64 `db:"kanji_id"`
	Onyomi  string `db:"onyomi"`
}

type Kunyomi struct {
	ID      uint64 `db:"id"`
	KanjiID uint64 `db:"kanji_id"`
	Kunyomi string `db:"kunyomi"`
}
