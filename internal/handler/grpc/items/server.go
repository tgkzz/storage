package items

import (
	"context"

	"github.com/tgkzz/storage/internal/models"
	"github.com/tgkzz/storage/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage1 "github.com/tgkzz/storage/gen/go/storage"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type serverApi struct {
	storage1.UnimplementedItemsServer
	service *service.Service
}

func Register(grpcServer *grpc.Server, service *service.Service) {
	storage1.RegisterItemsServer(grpcServer, &serverApi{service: service})
}

func (s *serverApi) CreateItem(ctx context.Context, request *storage1.CreateItemRequest) (*emptypb.Empty, error) {
	if request.Item == nil {
		return nil, status.Errorf(codes.InvalidArgument, "item is empty")
	}

	// maybe another ones?

	req := models.Item{
		Id:       int(request.GetItem().GetId()),
		Name:     request.GetItem().GetName(),
		Quantity: int(request.GetItem().GetQuantity()),
		Price: models.Price{
			CurrencyId: int(request.GetItem().GetPrice().GetCurrency()),
			Price:      float64(request.GetItem().GetPrice().GetPrice()),
		},
	}

	if err := s.service.Storage.CreateItem(ctx, req); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *serverApi) GetItemById(ctx context.Context, req *storage1.GetItemByIdRequest) (*storage1.GetItemByIdResponse, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id is empty")
	}

	res, err := s.service.Storage.GetItemById(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &storage1.GetItemByIdResponse{
		Item: &storage1.Item{
			Id:       int32(res.Id),
			Name:     res.Name,
			Quantity: int32(res.Quantity),
			Price: &storage1.Price{
				Currency: int32(res.Price.CurrencyId),
				Price:    float32(res.Price.Price),
			},
		},
	}, nil
}

func (s *serverApi) DeleteItemById(ctx context.Context, req *storage1.DeleteItemByIdRequest) (*emptypb.Empty, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id is empty")
	}

	if err := s.service.Storage.DeleteItemById(ctx, req.GetId()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
