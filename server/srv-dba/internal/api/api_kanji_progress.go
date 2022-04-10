package api

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
	m_kanji "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba"
)

type RepoKanjiProgress interface {
	GetKanjiProgressById(ctx context.Context, userID uint64, kanjiID uint64) (*m_kanji.KanjiProgress, error)
	ListKanjiProgressByTime(ctx context.Context, userID uint64, now time.Time) ([]*m_kanji.KanjiProgress, error)
	ListKanjiProgressByIDs(ctx context.Context, userID uint64, kanjiIDs []uint64) ([]*m_kanji.KanjiProgress, error)
	ListKanjiProgressBySRSLevel(ctx context.Context, userID uint64, srsLevel uint32) ([]*m_kanji.KanjiProgress, error)
}

func nullTimeToTimestamppb(t sql.NullTime) *timestamppb.Timestamp {
	if t.Valid {
		return timestamppb.New(t.Time)
	}
	return nil
}

func packKanjiProgress(progress *m_kanji.KanjiProgress) *pb.KanjiProgress {
	retKanjiProgress := &pb.KanjiProgress{
		Id:         progress.ID,
		Kanji:      packKanji(&progress.KanjiModel),
		SrsLevel:   progress.SRSLevel,
		UnlockDate: nullTimeToTimestamppb(progress.UnlockDate),
		NextDate:   nullTimeToTimestamppb(progress.NextDate),
		BurnDate:   nullTimeToTimestamppb(progress.BurnDate),
	}
	return retKanjiProgress
}

func (api *dbaAPI) GetKanjiProgressByIdV1(ctx context.Context, req *pb.GetKanjiProgressByIdV1Request,
) (*pb.GetKanjiProgressByIdV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "userID", req.GetUserId(), "kanjiID", req.GetKanjiId())

	kanjiProgress, err := api.repo.GetKanjiProgressById(ctx, req.GetUserId(), req.GetKanjiId())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if kanjiProgress == nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned nil from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "kanji progress not found")
	}
	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.GetKanjiProgressByIdV1Response{KanjiProgress: packKanjiProgress(kanjiProgress)}, nil

}

func (api *dbaAPI) ListKanjiProgressByTimeV1(ctx context.Context, req *pb.ListKanjiProgressByTimeV1Request,
) (*pb.ListKanjiProgressV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	now := time.Now()

	kanjiProgress, err := api.repo.ListKanjiProgressByTime(ctx, req.GetUserId(), now)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(kanjiProgress) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "kanji progress not found")
	}

	res := make([]*pb.KanjiProgress, len(kanjiProgress))
	for i, k := range kanjiProgress {
		res[i] = packKanjiProgress(k)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListKanjiProgressV1Response{KanjiProgress: res}, nil
}

func (api *dbaAPI) ListKanjiProgressByIdsV1(ctx context.Context, req *pb.ListKanjiProgressByIdsV1Request,
) (*pb.ListKanjiProgressV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	kanjiProgress, err := api.repo.ListKanjiProgressByIDs(ctx, req.GetUserId(), req.GetKanjiId())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(kanjiProgress) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "kanji progress not found")
	}

	res := make([]*pb.KanjiProgress, len(kanjiProgress))
	for i, k := range kanjiProgress {
		res[i] = packKanjiProgress(k)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListKanjiProgressV1Response{KanjiProgress: res}, nil
}

func (api *dbaAPI) ListKanjiProgressBySrsLevelV1(ctx context.Context, req *pb.ListKanjiProgressBySrsLevelV1Request,
) (*pb.ListKanjiProgressV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	kanjiProgress, err := api.repo.ListKanjiProgressBySRSLevel(ctx, req.GetUserId(), req.GetSrsLevel())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(kanjiProgress) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "kanji progress not found")
	}

	res := make([]*pb.KanjiProgress, len(kanjiProgress))
	for i, k := range kanjiProgress {
		res[i] = packKanjiProgress(k)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListKanjiProgressV1Response{KanjiProgress: res}, nil
}
