package models

type Sentence struct {
	ID           uint64                `db:"id"`
	WordID       uint64                `db:"word_id"`
	Origin       string                `db:"origin"`
	Translations []SentenceTranslation `db:"-"`
}

type SentenceTranslation struct {
	ID          uint64 `db:"id"`
	SentenceID  uint64 `db:"sentence_id"`
	Language    uint32 `db:"language"`
	Translation string `db:"translation"`
}
