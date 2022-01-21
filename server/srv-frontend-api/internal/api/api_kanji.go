package api

import (
	"context"

	"github.com/tomazis/kioku/server/srv-frontend-api/internal/logger"
	models "github.com/tomazis/kioku/server/srv-frontend-api/internal/models/kanji"
	pb "github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoKanji interface {
	GetKanji(ctx context.Context, kanjiID uint64) (*models.Kanji, error)
	ListKanji(ctx context.Context, level uint32) ([]*models.Kanji, error)
}

func (api *frontendAPI) GetKanjiV1(ctx context.Context, req *pb.GetKanjiV1Request,
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

func (api *frontendAPI) ListKanjiV1(ctx context.Context, req *pb.ListKanjiV1Request,
) (*pb.ListKanjiV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "ListKanjiV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	kanji, err := api.repo.ListKanji(ctx, req.GetLevel())
	if err != nil {
		logger.ErrorKV(ctx, "ListKanjiV1 -- failed to List from db", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(kanji) == 0 {
		logger.ErrorKV(ctx, "ListKanjiV1 -- returned zero items from db", "error", err)
		return nil, status.Error(codes.NotFound, "kanji not found")
	}

	res := make([]*pb.Kanji, len(kanji))
	for i, k := range kanji {
		res[i] = &pb.Kanji{
			Id:           k.ID,
			Kanji:        k.Kanji,
			Primary:      k.Primary,
			Level:        k.Level,
			Alternatives: k.Alternatives,
			Onyomi:       k.Onyomi,
			Kunyomi:      k.Kunyomi,
		}
	}

	logger.DebugKV(ctx, "ListKanjiV1 -- success")

	return &pb.ListKanjiV1Response{Kanji: res}, nil
}
