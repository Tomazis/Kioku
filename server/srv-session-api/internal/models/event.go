package models

import "time"

type Event struct {
	UserID        uint64
	UserLevel     uint32
	ProgressID    uint64
	ProgressLevel uint32
	SRSLevel      uint32
	Success       bool
	NextDate      *time.Time
	BurnDate      *time.Time
}

type Counter struct {
	KanjiSum    uint64
	UKanjiSum   uint64
	KanjiCount  uint64
	UKanjiCount uint64
	WordsSum    uint64
	UWordsSum   uint64
	WordsCount  uint64
	UWordsCount uint64
}
