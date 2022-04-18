package api

import (
	"context"
	"fmt"

	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
	m_helpers "github.com/tomazis/kioku/server/srv-dba/internal/models/helpers"
	pb "github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoHelpers interface {
	GetCounter(ctx context.Context, userID uint64, level uint32) (*m_helpers.Counter, error)
}

func (api *dbaAPI) GetCounterV1(ctx context.Context, req *pb.GetCounterV1Request,
) (*pb.GetCounterV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "userID", req.GetUserId(), "level", req.GetLevel())

	counter, err := api.repo.GetCounter(ctx, req.GetUserId(), req.GetLevel())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if counter == nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned nil from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "kanji not found")
	}
	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.GetCounterV1Response{
		Counter: &pb.Counter{
			KanjiCount:     counter.KanjiCount,
			UserKanjiCount: counter.UKanjiCount,
			WordsCount:     counter.WordsCount,
			UserWordsCount: counter.UWordsCount,
		},
	}, nil
}
