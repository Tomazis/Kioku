package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	gelf "github.com/snovichkov/zap-gelf"
	"github.com/tomazis/kioku/server/srv-dba/internal/api"
	"github.com/tomazis/kioku/server/srv-dba/internal/config"
	"github.com/tomazis/kioku/server/srv-dba/internal/database"
	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
	"github.com/tomazis/kioku/server/srv-dba/internal/repo"
	"github.com/tomazis/kioku/server/srv-dba/internal/server"
	"github.com/tomazis/kioku/server/srv-dba/internal/transfer"
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
		"importDB", cfg.Project.ImportDB,
	)

	migration := flag.Bool("migration", true, "Defines the migration start option")
	flag.Parse()

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(dsn, cfg.Database.Driver)
	if err != nil {
		logger.ErrorKV(ctx, "Failed init postgres", "error", err)
		return
	}
	defer db.Close()

	if *migration {
		if err = goose.Up(db.DB, cfg.Database.Migrations); err != nil {
			logger.ErrorKV(ctx, "Migration failed", "error", err)
			return
		}
	}

	if cfg.Project.ImportDB {
		if err = startTransfer(ctx, db); err != nil {
			return
		}
	}

	r := repo.NewRepo(db)

	dbaAPI := api.NewDbaAPI(r)

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

func startTransfer(ctx context.Context, db *sqlx.DB) (err error) {
	sqlitedb, err := database.NewSqlite("kanji.db", "sqlite3")
	if err != nil {
		logger.ErrorKV(ctx, "Failed init Sqlite", "error", err)
		return
	}
	defer sqlitedb.Close()

	logger.InfoKV(ctx, "Start transfer")

	err = transfer.Transfer(ctx, sqlitedb, db)

	if err != nil {
		logger.ErrorKV(ctx, "Error in transfer", "error", err)
		return err
	}
	return nil
}
