package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/tgkzz/storage/config"
	"github.com/tgkzz/storage/internal/app"
	pkgLogger "github.com/tgkzz/storage/pkg/logger"
)

const (
	cfgPath = "CONFIG_PATH"
	env     = "ENV"
)

func AppRun() {
	defer func() {
		if r := recover(); r != nil {
			slog.Info("Recovered Error after panic", slog.Any("panic", r))
		}
	}()
	cPath := os.Getenv(cfgPath)
	cfg := config.MustRead(cPath)

	var logger *slog.Logger
	switch cfg.Env != "" {
	case true:
		logger = pkgLogger.SetupLogger(env)
	default:
		logger = pkgLogger.SetupLogger("local")
	}

	c := context.Background()

	ctx, stop := signal.NotifyContext(c, os.Interrupt)
	defer stop()

	a := app.New(ctx, logger, cfg)

	go func() {
		a.ItemGRPCServer.MustRun()
	}()

	go func() {
		a.StorageGRPCServer.MustRun()
	}()

	<-ctx.Done()

	go func() {
		a.ItemGRPCServer.Stop()
	}()
	go func() {
		a.StorageGRPCServer.Stop()
	}()
}
