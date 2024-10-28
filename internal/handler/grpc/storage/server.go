package storage

import (
	"context"
	"github.com/tgkzz/storage/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	if req.Items == nil {
		return nil, status.Errorf(codes.InvalidArgument, "items is required")
	}

	if req.Username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username is required")
	}

	var r []models.Item
	for _, item := range req.Items {
		r = append(r, models.Item{
			Id:       int(item.GetId()),
			Name:     item.GetName(),
			Quantity: int(item.Quantity),
			Price: models.Price{
				CurrencyId: int(item.Price.Currency),
				Price:      float64(item.Price.Price),
			},
		})
	}

	if err := s.service.Storage.CreateOrder(ctx, r, req.GetUsername()); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &storage1.CreateOrderResponse{}, nil
}

func (s *serverApi) CancelOrder(ctx context.Context, req *storage1.CancelOrderRequest) (*storage1.CancelOrderResponse, error) {
	return &storage1.CancelOrderResponse{}, nil
}
