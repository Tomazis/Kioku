package repo

import (
	"context"
	"errors"

	models "github.com/tomazis/kioku/server/srv-frontend-api/internal/models/word"
)

func (r *repo) GetWord(ctx context.Context, wordID uint64) (*models.Word, error) {
	return nil, errors.New("not implemented")
}
func (r *repo) ListWordsByLevel(ctx context.Context, level uint32) ([]*models.Word, error) {
	return nil, errors.New("not implemented")
}
func (r *repo) ListWordByKanji(ctx context.Context, kanjiID uint64) ([]*models.Word, error) {
	return nil, errors.New("not implemented")
}
