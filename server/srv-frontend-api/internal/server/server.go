package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tomazis/kioku/server/srv-frontend-api/internal/api"
	"github.com/tomazis/kioku/server/srv-frontend-api/internal/config"
	"github.com/tomazis/kioku/server/srv-frontend-api/internal/logger"
	"github.com/tomazis/kioku/server/srv-frontend-api/internal/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	pb "github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api"
)

type gRPCServer struct{}

func NewGRPCServer() *gRPCServer {
	return &gRPCServer{}
}

func (s *gRPCServer) Start(ctx context.Context, cfg *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcAddr := fmt.Sprintf("%s:%v", cfg.Grpc.Host, cfg.Grpc.Port)

	listen, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on grpc: %w", err)
	}
	defer listen.Close()

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(cfg.Grpc.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(cfg.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(cfg.Grpc.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(cfg.Grpc.Timeout) * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
		)),
	)

	r := repo.NewRepo(time.Duration(cfg.GrpcDBA.Timeout)*time.Second, fmt.Sprintf("%s:%d", cfg.GrpcDBA.Host, cfg.GrpcDBA.Port))

	pb.RegisterSrvFrontendApiServiceServer(grpcServer, api.NewFrontendAPI(r))

	go func() {
		logger.InfoKV(ctx, "gRPC Server is listening", "address", grpcAddr)
		if err := grpcServer.Serve(listen); err != nil {
			logger.FatalKV(ctx, "Server exited with error", "error", err)
		}
	}()

	if cfg.Project.Debug {
		reflection.Register(grpcServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-quit:
		logger.InfoKV(ctx, "signal.Notify received", "signal", sig)
	case done := <-ctx.Done():
		logger.InfoKV(ctx, "ctx.Done received", "ctx", done)
	}

	grpcServer.GracefulStop()
	logger.InfoKV(ctx, "gRPC Server is stopped")

	return nil
}
