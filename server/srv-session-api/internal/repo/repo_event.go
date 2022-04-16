package repo

import "github.com/tomazis/kioku/server/srv-session-api/internal/models"

// Get User Level
// If kanji/word level == user level then check if both (kanji and word) progress% > 75%
// If true then user level +1 and add new kanji and words to progress corresponding to level
func (r *repo) ProceedEvent(e *models.Event) error {
	//Get user level
	userLevel, err := getUserLevel(e.UserID)
	if err != nil {
		return err
	}
	if e.ProgressLevel >= userLevel {
		progSum, levelSum, err := getSumProgress(e.UserID, e.UserLevel)
		if err != nil {
			return err
		}
		if float64(progSum)/float64(levelSum) >= 0.75 {

		}
	}
}

func getUserLevel(userID uint64) (uint32, error) {
	return 0, nil
}

func getSumProgress(userID uint64, userLevel uint32) (uint64, uint64, error) {
	return 0, 0, nil
}
