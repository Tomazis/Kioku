package api

import (
	"context"

	m_kanji "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
)

type RepoKanji interface {
	GetKanji(ctx context.Context, kanjiID uint64) (*m_kanji.Kanji, error)
	ListKanji(ctx context.Context, level uint32) ([]*m_kanji.Kanji, error)
}
