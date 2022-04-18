package repo

import (
	"context"
	"errors"
	"time"

	"github.com/tomazis/kioku/server/srv-session-api/internal/models"
)

//Srs level 0 for study

// Get User Level
// If kanji/word level == user level then check if both (kanji and word) progress% > 85%
// If true then user level +1 and add new kanji and words to progress corresponding to level
func (r *repo) ProceedEvent(ctx context.Context, e *models.Event, isKanji bool) error {
	counter, err := getCounter(ctx, e.UserID, e.UserLevel)
	if err != nil {
		return err
	}

	err = setProgressTime(e)
	if err != nil {
		return err
	}

	err = updateProgress(ctx, e)
	if err != nil {
		return err
	}

	if isKanji && counter.KanjiCount != counter.UKanjiCount {
		err = addNewKanjiAndUnlockWords(ctx, e.UserID, e.ProgressLevel)
		if err != nil {
			return err
		}
	}

	if e.ProgressLevel == e.UserLevel {
		completion := float64(counter.UWordsCount+counter.UKanjiCount) / float64(counter.KanjiCount+counter.WordsCount) * 100

		if completion >= 85 {
			userLevelUp(ctx, e.UserID)
		}
	}
	return nil
}

func setProgressTime(e *models.Event) error {
	if e.SRSLevel == 9 {
		return errors.New("Already burned")
	}

	if e.Success {
		e.SRSLevel++
	} else if e.SRSLevel > 1 {
		e.SRSLevel--
	}

	now := time.Now()

	e.BurnDate = nil
	e.NextDate = nil

	switch e.SRSLevel {
	case 1:
		now = now.Add(time.Hour * 4)
		e.NextDate = &now
	case 2:
		now = now.Add(time.Hour * 8)
		e.NextDate = &now
	case 3:
		now = now.AddDate(0, 0, 1)
		e.NextDate = &now
	case 4:
		now = now.AddDate(0, 0, 2)
		e.NextDate = &now
	case 5:
		now = now.AddDate(0, 0, 7)
		e.NextDate = &now
	case 6:
		now = now.AddDate(0, 0, 14)
		e.NextDate = &now
	case 7:
		now = now.AddDate(0, 1, 0)
		e.NextDate = &now
	case 8:
		now = now.AddDate(0, 4, 0)
		e.NextDate = &now
	case 9:
		e.BurnDate = &now
	}
	return nil
}

func getCounter(ctx context.Context, userID uint64, userLevel uint32) (*models.Counter, error) {
	var counter *models.Counter
	return counter, nil
}

//add 50% kanji to study and words that hasn't had that level kanji
func userLevelUp(ctx context.Context, userID uint64) error {
	return nil
}

func addNewKanjiAndUnlockWords(ctx context.Context, userID uint64, progressLevel uint32) error {
	return nil
}

func updateProgress(ctx context.Context, e *models.Event) error {
	return nil
}
