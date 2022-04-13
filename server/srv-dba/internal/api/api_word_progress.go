package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
	m_word "github.com/tomazis/kioku/server/srv-dba/internal/models/word"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba"
)

type RepoWordProgress interface {
	GetWordProgressById(ctx context.Context, userID uint64, wordID uint64) (*m_word.WordProgress, error)
	ListWordProgressByTime(ctx context.Context, userID uint64, now time.Time, limit uint64, offset uint64) ([]*m_word.WordProgress, error)
	ListWordProgressByIDs(ctx context.Context, userID uint64, wordIDs []uint64) ([]*m_word.WordProgress, error)
	ListWordProgressBySRSLevel(ctx context.Context, userID uint64, srsLevel uint32, limit uint64, offset uint64) ([]*m_word.WordProgress, error)
	AddWordProgress(ctx context.Context, userID uint64, wordID []uint64) (bool, error)
	UpdateWordProgress(ctx context.Context, progressID uint64, srsLevel uint32, nextDate *time.Time, burnDate *time.Time) (bool, error)
}

func packWordProgress(progress *m_word.WordProgress) *pb.WordProgress {
	retWordProgress := &pb.WordProgress{
		Id:         progress.ID,
		Word:       packWord(&progress.WordModel),
		SrsLevel:   progress.SRSLevel,
		UnlockDate: nullTimeToTimestamppb(progress.UnlockDate),
		NextDate:   nullTimeToTimestamppb(progress.NextDate),
		BurnDate:   nullTimeToTimestamppb(progress.BurnDate),
	}
	return retWordProgress
}

func (api *dbaAPI) GetWordProgressByIdV1(ctx context.Context, req *pb.GetWordProgressByIdV1Request,
) (*pb.GetWordProgressByIdV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "userID", req.GetUserId(), "wordID", req.GetWordId())

	wordProgress, err := api.repo.GetWordProgressById(ctx, req.GetUserId(), req.GetWordId())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if wordProgress == nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned nil from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "word progress not found")
	}
	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.GetWordProgressByIdV1Response{WordProgress: packWordProgress(wordProgress)}, nil

}

func (api *dbaAPI) ListWordProgressByTimeV1(ctx context.Context, req *pb.ListWordProgressByTimeV1Request,
) (*pb.ListWordProgressV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	now := time.Now()

	wordProgress, err := api.repo.ListWordProgressByTime(ctx, req.GetUserId(), now, req.GetLimit(), req.GetOffset())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(wordProgress) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "word progress not found")
	}

	res := make([]*pb.WordProgress, len(wordProgress))
	for i, k := range wordProgress {
		res[i] = packWordProgress(k)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListWordProgressV1Response{WordProgress: res}, nil
}

func (api *dbaAPI) ListWordProgressByIdsV1(ctx context.Context, req *pb.ListWordProgressByIdsV1Request,
) (*pb.ListWordProgressV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	wordProgress, err := api.repo.ListWordProgressByIDs(ctx, req.GetUserId(), req.GetWordId())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(wordProgress) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "word progress not found")
	}

	res := make([]*pb.WordProgress, len(wordProgress))
	for i, k := range wordProgress {
		res[i] = packWordProgress(k)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListWordProgressV1Response{WordProgress: res}, nil
}

func (api *dbaAPI) ListWordProgressBySrsLevelV1(ctx context.Context, req *pb.ListWordProgressBySrsLevelV1Request,
) (*pb.ListWordProgressV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	wordProgress, err := api.repo.ListWordProgressBySRSLevel(ctx, req.GetUserId(), req.GetSrsLevel(), req.GetLimit(), req.GetOffset())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(wordProgress) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "word progress not found")
	}

	res := make([]*pb.WordProgress, len(wordProgress))
	for i, k := range wordProgress {
		res[i] = packWordProgress(k)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListWordProgressV1Response{WordProgress: res}, nil
}

func (api *dbaAPI) AddWordProgressV1(ctx context.Context, req *pb.AddWordProgressV1Request,
) (*pb.DefaultWordProgressV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "userID", req.GetUserId(), "wordID", req.GetWordId())

	wordProgress, err := api.repo.AddWordProgress(ctx, req.GetUserId(), req.GetWordId())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.DefaultWordProgressV1Response{Success: wordProgress}, nil
}

func (api *dbaAPI) UpdateWordProgressV1(ctx context.Context, req *pb.UpdateWordProgressV1Request,
) (*pb.DefaultWordProgressV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if req.GetSrsLevel() == 9 && (req.GetNextDate() != nil || req.GetBurnDate() == nil) {
		err := errors.New("SRS Level 9 but Next Date is not nil and Burn Date is nil")
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if req.GetSrsLevel() < 9 && (req.GetBurnDate() != nil || req.GetNextDate() == nil) {
		err := errors.New("SRS Level < 9 but Burn Date is not nil and Next Date is nil")
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "progressID", req.GetProgressId(),
		"srs_level", req.GetSrsLevel(), "next_date", req.GetNextDate(), "burn_date", req.GetBurnDate())
	var t time.Time
	var nextDate *time.Time

	if req.GetNextDate() != nil {
		t = req.GetNextDate().AsTime()
		nextDate = &t
	}
	var burnDate *time.Time
	if req.GetBurnDate() != nil {
		t = req.GetBurnDate().AsTime()
		burnDate = &t
	}

	wordProgress, err := api.repo.UpdateWordProgress(ctx, req.GetProgressId(), req.GetSrsLevel(), nextDate, burnDate)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.DefaultWordProgressV1Response{Success: wordProgress}, nil
}
