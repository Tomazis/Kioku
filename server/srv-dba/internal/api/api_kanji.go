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
	GetKanjiByID(ctx context.Context, kanjiID uint64) (*m_kanji.Kanji, error)
	ListKanjiByLevel(ctx context.Context, level uint32) ([]*m_kanji.Kanji, error)
	ListKanjiByIDs(ctx context.Context, ids []uint64) ([]*m_kanji.Kanji, error)
}

func (api *dbaAPI) GetKanjiByIdV1(ctx context.Context, req *pb.GetKanjiByIdV1Request,
) (*pb.GetKanjiByIdV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "GetKanjiV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "kanjiID", req.GetKanjiId())

	kanji, err := api.repo.GetKanjiByID(ctx, req.GetKanjiId())
	if err != nil {
		logger.ErrorKV(ctx, "GetKanjiV1 -- failed to get from db", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if kanji == nil {
		logger.ErrorKV(ctx, "GetKanjiV1 -- returned nil from db", "error", err)
		return nil, status.Error(codes.NotFound, "kanji not found")
	}
	logger.DebugKV(ctx, "GetKanjiV1 -- success")

	return &pb.GetKanjiByIdV1Response{Kanji: &pb.Kanji{
		Id:           kanji.ID,
		Kanji:        kanji.Kanji,
		Primary:      kanji.Primary,
		Level:        kanji.Level,
		Alternatives: aggStringToSlice(kanji.Alternatives, "|"),
		Onyomi:       aggStringToSlice(kanji.Onyomi, "|"),
		Kunyomi:      aggStringToSlice(kanji.Kunyomi, "|"),
	}}, nil
}

func (api *dbaAPI) ListKanjiByLevelV1(ctx context.Context, req *pb.ListKanjiByLevelV1Request,
) (*pb.ListKanjiV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "ListKanjiByLevelV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	kanji, err := api.repo.ListKanjiByLevel(ctx, req.GetLevel())
	if err != nil {
		logger.ErrorKV(ctx, "ListKanjiByLevelV1 -- failed to List from db", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(kanji) == 0 {
		logger.ErrorKV(ctx, "ListKanjiByLevelV1 -- returned zero items from db", "error", err)
		return nil, status.Error(codes.NotFound, "kanji not found")
	}

	res := make([]*pb.Kanji, len(kanji))
	for i, k := range kanji {
		res[i] = &pb.Kanji{
			Id:           k.ID,
			Kanji:        k.Kanji,
			Primary:      k.Primary,
			Level:        k.Level,
			Alternatives: aggStringToSlice(k.Alternatives, "|"),
			Onyomi:       aggStringToSlice(k.Onyomi, "|"),
			Kunyomi:      aggStringToSlice(k.Kunyomi, "|"),
		}
	}

	logger.DebugKV(ctx, "ListKanjiByLevelV1 -- success")

	return &pb.ListKanjiV1Response{Kanji: res}, nil
}

func (api *dbaAPI) ListKanjiByIdsV1(ctx context.Context, req *pb.ListKanjiByIdsV1Request,
) (*pb.ListKanjiV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "ListKanjiByIdsV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	kanji, err := api.repo.ListKanjiByIDs(ctx, req.GetKanjiId())
	if err != nil {
		logger.ErrorKV(ctx, "ListKanjiByIdsV1 -- failed to List from db", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(kanji) == 0 {
		logger.ErrorKV(ctx, "ListKanjiByIdsV1 -- returned zero items from db", "error", err)
		return nil, status.Error(codes.NotFound, "kanji not found")
	}

	res := make([]*pb.Kanji, len(kanji))
	for i, k := range kanji {
		res[i] = &pb.Kanji{
			Id:           k.ID,
			Kanji:        k.Kanji,
			Primary:      k.Primary,
			Level:        k.Level,
			Alternatives: aggStringToSlice(k.Alternatives, "|"),
			Onyomi:       aggStringToSlice(k.Onyomi, "|"),
			Kunyomi:      aggStringToSlice(k.Kunyomi, "|"),
		}
	}

	logger.DebugKV(ctx, "ListKanjiByIdsV1 -- success")

	return &pb.ListKanjiV1Response{Kanji: res}, nil
}
