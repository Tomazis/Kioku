package api

import (
	"context"
	"fmt"

	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
	m_kanji "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba"
)

type RepoKanji interface {
	GetKanjiByID(ctx context.Context, kanjiID uint64) (*m_kanji.Kanji, error)
	ListKanjiByLevel(ctx context.Context, level uint32, limit uint64, offset uint64, min bool) ([]*m_kanji.Kanji, error)
	ListKanjiByIDs(ctx context.Context, ids []uint64, min bool) ([]*m_kanji.Kanji, error)
}

func packKanji(kanji *m_kanji.Kanji) *pb.Kanji {
	retKanji := &pb.Kanji{
		Id:           kanji.ID,
		Kanji:        kanji.Kanji,
		Primary:      kanji.Primary,
		Level:        kanji.Level,
		Alternatives: aggStringToSlice(kanji.Alternatives, "|"),
		Onyomi:       aggStringToSlice(kanji.Onyomi, "|"),
		Kunyomi:      aggStringToSlice(kanji.Kunyomi, "|"),
	}
	return retKanji
}

func (api *dbaAPI) GetKanjiByIdV1(ctx context.Context, req *pb.GetKanjiByIdV1Request,
) (*pb.GetKanjiByIdV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "kanjiID", req.GetKanjiId())

	kanji, err := api.repo.GetKanjiByID(ctx, req.GetKanjiId())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if kanji == nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned nil from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "kanji not found")
	}
	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.GetKanjiByIdV1Response{Kanji: packKanji(kanji)}, nil
}

func (api *dbaAPI) ListKanjiByLevelV1(ctx context.Context, req *pb.ListKanjiByLevelV1Request,
) (*pb.ListKanjiV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	kanji, err := api.repo.ListKanjiByLevel(ctx, req.GetLevel(), req.GetLimit(), req.GetOffset(), req.GetMin())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(kanji) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "kanji not found")
	}

	res := make([]*pb.Kanji, len(kanji))
	for i, k := range kanji {
		res[i] = packKanji(k)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListKanjiV1Response{Kanji: res}, nil
}

func (api *dbaAPI) ListKanjiByIdsV1(ctx context.Context, req *pb.ListKanjiByIdsV1Request,
) (*pb.ListKanjiV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	kanji, err := api.repo.ListKanjiByIDs(ctx, req.GetKanjiId(), req.GetMin())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(kanji) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "kanji not found")
	}

	res := make([]*pb.Kanji, len(kanji))
	for i, k := range kanji {
		res[i] = packKanji(k)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListKanjiV1Response{Kanji: res}, nil
}
