package app

import (
	"context"
	"log/slog"

	"github.com/tgkzz/storage/config"
	"github.com/tgkzz/storage/internal/app/grpc"
	"github.com/tgkzz/storage/internal/repository"
	"github.com/tgkzz/storage/internal/service"
)

type App struct {
	ItemGRPCServer    *grpc.App
	StorageGRPCServer *grpc.App
}

func New(ctx context.Context, log *slog.Logger, cfg *config.Config) *App {
	repo, err := repository.NewPostgresRepository(ctx, cfg.PostgresDatabase.Url)
	if err != nil {
		panic(err)
	}

	srv := service.NewService(log, repo, repo)

	grpcItemServer := grpc.New(log, cfg.GrpcItemServer.Port, srv, grpc.ItemService)

	grpcStorageServer := grpc.New(log, cfg.GrpcStorageServer.Port, srv, grpc.StorageService)

	return &App{
		ItemGRPCServer:    grpcItemServer,
		StorageGRPCServer: grpcStorageServer,
	}
}
