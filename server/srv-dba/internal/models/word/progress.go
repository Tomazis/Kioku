package models

import "time"

type WordProgress struct {
	ID         uint64    `db:"id"`
	WordID     uint64    `db:"word_id"`
	UserID     uint64    `db:"user_id"`
	SRSLevel   uint32    `db:"srs_level"`
	UnlockDate time.Time `db:"unlock_date"`
	NextDate   time.Time `db:"next_date"`
	BurnDate   time.Time `db:"burn_date"`
}
