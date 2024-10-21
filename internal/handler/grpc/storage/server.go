package storage

import (
	"context"

	"github.com/tgkzz/storage/internal/service"

	storage1 "github.com/tgkzz/storage/gen/go/storage"
	"google.golang.org/grpc"
)

type serverApi struct {
	storage1.UnimplementedStorageServer
	service *service.Service
}

func Register(server *grpc.Server, service *service.Service) {
	storage1.RegisterStorageServer(server, &serverApi{service: service})
}

func (s *serverApi) CreateOrder(ctx context.Context, req *storage1.CreateOrderRequest) (*storage1.CreateOrderResponse, error) {
	return &storage1.CreateOrderResponse{}, nil
}

func (s *serverApi) CancelOrder(ctx context.Context, req *storage1.CancelOrderRequest) (*storage1.CancelOrderResponse, error) {
	return &storage1.CancelOrderResponse{}, nil
}
