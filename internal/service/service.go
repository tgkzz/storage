package service

import (
	"context"
	"log/slog"

	"github.com/tgkzz/storage/internal/repository"
	"github.com/tgkzz/storage/internal/service/storage"

	"github.com/tgkzz/storage/internal/models"
)

type Service struct {
	Storage StorageService
}

type StorageService interface {
	CreateItem(ctx context.Context, item models.Item) error
	DeleteItemById(ctx context.Context, id string) error
	GetItemById(ctx context.Context, id string) (models.Item, error)
	UpdatePriceByItemId(ctx context.Context, id string, newPrice float64, currencyId string) error
	UpdateQuantityById(ctx context.Context, id string, newQuantity int) error

	CreateOrder(ctx context.Context, items []models.Item, username string) error
}

func NewService(logger *slog.Logger,
	itemRepo repository.ItemRepository, currRepo repository.CurrencyRepository,
) *Service {
	return &Service{
		Storage: storage.NewStorageService(logger, itemRepo, currRepo),
	}
}
