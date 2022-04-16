package main

import (
	"context"
	"fmt"
	"os"
	"time"

	gelf "github.com/snovichkov/zap-gelf"
	"github.com/tomazis/kioku/server/srv-session-api/internal/api"
	"github.com/tomazis/kioku/server/srv-session-api/internal/config"
	"github.com/tomazis/kioku/server/srv-session-api/internal/logger"
	"github.com/tomazis/kioku/server/srv-session-api/internal/repo"
	"github.com/tomazis/kioku/server/srv-session-api/internal/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yaml"); err != nil {
		panic(err)
	}
	cfg := config.GetConfigInstance()

	syncLogger := initLogger(ctx, cfg)
	defer syncLogger()

	logger.InfoKV(ctx, "Start",
		"Name", cfg.Project.Name,
		"version", cfg.Project.Version,
		"commitHash", cfg.Project.CommitHash,
		"debug", cfg.Project.Debug,
		"environment", cfg.Project.Environment,
	)

	r := repo.NewRepo(time.Duration(cfg.GrpcDBA.Timeout)*time.Second, fmt.Sprintf("%s:%d", cfg.GrpcDBA.Host, cfg.GrpcDBA.Port))

	dbaAPI := api.NewSessionAPI(r)

	if err := server.NewGRPCServer().Start(ctx, &cfg, &dbaAPI); err != nil {
		logger.ErrorKV(ctx, "Failed to start gRPC server", "error", err)
		return
	}
}

func initLogger(ctx context.Context, cfg config.Config) (syncFn func()) {
	loggingLevel := zap.InfoLevel
	if cfg.Project.Debug {
		loggingLevel = zap.DebugLevel
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	gelfCore, err := gelf.NewCore(
		gelf.Addr(cfg.Telemetry.GraylogPath),
		gelf.Level(loggingLevel),
	)
	if err != nil {
		logger.FatalKV(ctx, "sql.Open() error", "err", err)
	}

	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore, gelfCore))

	sugaredLogger := notSugaredLogger.Sugar()
	logger.SetLogger(sugaredLogger.With(
		"service", cfg.Project.ServiceName,
	))

	return func() {
		err := notSugaredLogger.Sync()
		if err != nil {
			logger.FatalKV(ctx, "not sugared logger sync error", "err", err)
		}
	}

}
