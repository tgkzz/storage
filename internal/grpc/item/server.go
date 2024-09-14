package item

import (
	"context"
	"fmt"
	storage1 "github.com/tgkzz/storage/gen/go/storage"
	"github.com/tgkzz/storage/internal/domain/models"
	"github.com/tgkzz/storage/internal/service/item"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverApi struct {
	storage1.UnimplementedItemServer
	item item.ItemService
}

func Register(grpcServer *grpc.Server, itemService item.ItemService) {
	storage1.RegisterItemServer(grpcServer, &serverApi{item: itemService})
}

func (s *serverApi) GetItemById(ctx context.Context, req *storage1.GetItemRequest) (*storage1.GetItemResponse, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id of item must be given")
	}

	resp, err := s.item.GetItemById(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get item")
	}

	return &storage1.GetItemResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Desc,
		Quantity:    resp.Quantity,
		Price:       resp.Price,
		Currency:    resp.Currency,
	}, err
}

func (s *serverApi) DeleteItem(ctx context.Context, req *storage1.DeleteItemRequest) (*storage1.DeleteItemResponse, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id of item must be given")
	}

	if err := s.item.DeleteItemById(ctx, req.GetId()); err != nil {
		return nil, status.Error(codes.Internal, "failed to delete item")
	}

	return &storage1.DeleteItemResponse{
		Success: true,
		Message: fmt.Sprintf("succesfully deleted product with id %d", req.Id),
	}, nil
}

func (s *serverApi) UpdateItem(ctx context.Context, req *storage1.UpdateItemRequest) (*storage1.UpdateItemResponse, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id of item must be given")
	}

	var opts []models.OptionFunc

	if req.Name != nil {
		opts = append(opts, models.SetName(req.GetName()))
	}

	if req.Description != nil {
		opts = append(opts, models.SetDesc(req.GetDescription()))
	}

	if req.Quantity != nil {
		opts = append(opts, models.SetQuantity(req.GetQuantity()))
	}

	if req.Price != nil {
		opts = append(opts, models.SetPrice(req.GetPrice()))
	}

	if req.Currency != nil {
		opts = append(opts, models.SetCurrency(req.GetCurrency()))
	}

	if len(opts) == 0 {
		return nil, status.Error(codes.OK, "nothing to update")
	}

	if err := s.item.UpdateItem(ctx, req.GetId(), opts...); err != nil {
		return nil, status.Error(codes.Internal, "failed to update item")
	}

	return &storage1.UpdateItemResponse{
		Success: true,
		Message: fmt.Sprintf("succesfully update product with id %d", req.GetId()),
	}, nil
}

func (s *serverApi) CreateItem(ctx context.Context, req *storage1.CreateItemRequest) (*storage1.CreateItemResponse, error) {
	return nil, nil
}
