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
