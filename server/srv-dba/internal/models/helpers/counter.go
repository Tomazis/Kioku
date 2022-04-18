package models

type Counter struct {
	KanjiCount  uint64 `db:"kanji_count"`
	UKanjiCount uint64 `db:"user_kanji_count"`
	WordsCount  uint64 `db:"words_count"`
	UWordsCount uint64 `db:"user_words_count"`
}
