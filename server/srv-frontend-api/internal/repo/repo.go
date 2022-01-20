package repo

import (
	"context"
	"errors"

	models "github.com/tomazis/kioku/server/srv-frontend-api/internal/models/kanji"
)

type repo struct{}

func NewRepo() *repo {
	return &repo{}
}

func (r *repo) Get(ctx context.Context, kanjiID uint64) (*models.Kanji, error) {
	return nil, errors.New("not implemented")
}

func (r *repo) List(ctx context.Context, level uint32) ([]models.Kanji, error) {
	return nil, errors.New("not implemented")
}
