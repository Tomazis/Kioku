package api

import (
	"context"
	"fmt"

	"github.com/tomazis/kioku/server/srv-session-api/internal/logger"
	"github.com/tomazis/kioku/server/srv-session-api/internal/models"
	pb "github.com/tomazis/kioku/server/srv-session-api/pkg/srv-session-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoEvent interface {
	ProceedEvent(ctx context.Context, e *models.Event, isKanji bool) error
}

func (api *sessionAPI) KanjiEventV1(ctx context.Context, req *pb.KanjiEventV1Request,
) (*pb.EventV1Reposnse, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	e := models.Event{
		UserID:        req.GetEvent().GetUserId(),
		UserLevel:     req.GetEvent().GetUserLevel(),
		ProgressID:    req.GetEvent().GetProgressId(),
		ProgressLevel: req.GetEvent().GetProgressLevel(),
		SRSLevel:      req.GetEvent().GetSrsLevel(),
		Success:       req.GetEvent().GetSuccess(),
		NextDate:      nil,
		BurnDate:      nil,
	}

	err := api.repo.ProceedEvent(ctx, &e, true)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to update", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.EventV1Reposnse{Success: true}, nil
}

func (api *sessionAPI) WordEventV1(ctx context.Context, req *pb.WordEventV1Request,
) (*pb.EventV1Reposnse, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	e := models.Event{
		UserID:        req.GetEvent().GetUserId(),
		UserLevel:     uint32(req.GetEvent().GetUserId()),
		ProgressID:    req.GetEvent().ProgressId,
		ProgressLevel: req.GetEvent().GetProgressLevel(),
		SRSLevel:      req.GetEvent().GetSrsLevel(),
		Success:       req.GetEvent().GetSuccess(),
		NextDate:      nil,
		BurnDate:      nil,
	}

	err := api.repo.ProceedEvent(ctx, &e, true)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to update", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.EventV1Reposnse{Success: true}, nil
}
