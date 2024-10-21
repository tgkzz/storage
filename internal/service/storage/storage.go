package storage

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/tgkzz/storage/internal/models"
	"github.com/tgkzz/storage/internal/repository"
)

type StorageService struct {
	logger   *slog.Logger
	itemRepo repository.ItemRepository
	currRepo repository.CurrencyRepository
}

func NewStorageService(logger *slog.Logger,
	itemRepo repository.ItemRepository, currRepo repository.CurrencyRepository,
) *StorageService {
	return &StorageService{
		logger:   logger,
		itemRepo: itemRepo,
		currRepo: currRepo,
	}
}

func (s *StorageService) CreateItem(ctx context.Context, item models.Item) error {
	const op = "storageService.CreateItem"

	log := s.logger.With(
		slog.String("op", op),
		slog.Any("item", item),
	)

	currId := strconv.Itoa(item.Price.CurrencyId)

	currency, err := s.currRepo.GetCurrencyById(ctx, currId)
	if err != nil {
		log.Error("currency not found", slog.String("currencyId", currId))
		return err
	}

	log.Info("currency found", slog.String("currencyCode", currency.Code))

	if err = s.itemRepo.InsertItem(ctx, &item); err != nil {
		log.Error("error inserting item", slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (s *StorageService) DeleteItemById(ctx context.Context, id string) error {
	const op = "storageService.DeleteItemById"

	log := s.logger.With(
		slog.String("op", op),
		slog.String("id", id),
	)

	if err := s.itemRepo.DeleteItemById(ctx, id); err != nil {
		log.Error("error deleting item", slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (s *StorageService) GetItemById(ctx context.Context, id string) (models.Item, error) {
	const op = "storageService.GetItemById"

	log := s.logger.With(
		slog.String("op", op),
		slog.String("id", id),
	)

	item, err := s.itemRepo.GetItemById(ctx, id)
	if err != nil {
		log.Error("error getting item", slog.String("error", err.Error()))
		return models.Item{}, err
	}

	return *item, nil
}

func (s *StorageService) UpdatePriceByItemId(ctx context.Context, id string, newPrice float64, currencyId string) error {
	const op = "storageService.UpdatePriceByItemId"

	log := s.logger.With(
		slog.String("op", op),
		slog.String("itemId", id),
		slog.String("currencyId", currencyId),
	)

	item, err := s.GetItemById(ctx, id)
	if err != nil {
		return err
	}

	item.Price.Price = newPrice
	item.Price.CurrencyId, err = strconv.Atoi(currencyId)
	if err != nil {
		return err
	}

	if err = s.CreateItem(ctx, item); err != nil {
		log.Error("error creating item", slog.String("error", err.Error()))
		return err
	}

	s.logger.Info("item price successfully updated", slog.String("itemId", strconv.Itoa(item.Id)))

	return nil
}

func (s *StorageService) UpdateQuantityById(id string, newQuantity int) error {
	// TODO: must be realized using update method

	return nil
}
