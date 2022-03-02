package repo

import (
	"context"
	"errors"

	models "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
)

func (r *repo) GetKanji(ctx context.Context, kanjiID uint64) (*models.Kanji, error) {
	return nil, errors.New("not implemented")
}

func (r *repo) ListKanji(ctx context.Context, level uint32) ([]*models.Kanji, error) {
	return nil, errors.New("not implemented")
}
