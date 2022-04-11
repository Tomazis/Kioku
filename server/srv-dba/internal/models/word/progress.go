package models

import "database/sql"

type WordProgress struct {
	ID         uint64       `db:"id"`
	WordID     uint64       `db:"word_id"`
	UserID     uint64       `db:"user_id"`
	SRSLevel   uint32       `db:"srs_level"`
	UnlockDate sql.NullTime `db:"unlock_date"`
	NextDate   sql.NullTime `db:"next_date"`
	BurnDate   sql.NullTime `db:"burn_date"`
	WordModel  Word         `db:"-"`
}
