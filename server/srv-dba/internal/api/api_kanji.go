package api

import (
	"context"

	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
	m_kanji "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba"
)

type RepoKanji interface {
	GetKanji(ctx context.Context, kanjiID uint64) (*m_kanji.Kanji, error)
	ListKanji(ctx context.Context, level uint32) ([]*m_kanji.Kanji, error)
}

func (api *dbaAPI) GetKanjiV1(ctx context.Context, req *pb.GetKanjiV1Request,
) (*pb.GetKanjiV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "GetKanjiV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "kanjiID", req.GetKanjiId())

	kanji, err := api.repo.GetKanji(ctx, req.GetKanjiId())
	if err != nil {
		logger.ErrorKV(ctx, "GetKanjiV1 -- failed to get from db", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if kanji == nil {
		logger.ErrorKV(ctx, "GetKanjiV1 -- returned nil from db", "error", err)
		return nil, status.Error(codes.NotFound, "kanji not found")
	}
	logger.DebugKV(ctx, "GetKanjiV1 -- success")

	return &pb.GetKanjiV1Response{Kanji: &pb.Kanji{
		Id:           kanji.ID,
		Kanji:        kanji.Kanji,
		Primary:      kanji.Primary,
		Level:        kanji.Level,
		Alternatives: kanji.Alternatives,
		Onyomi:       kanji.Onyomi,
		Kunyomi:      kanji.Kunyomi,
	}}, nil
}
