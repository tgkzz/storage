package item

import (
	"context"
	"errors"
	"fmt"
	"github.com/tgkzz/storage/internal/domain/dto"
	"github.com/tgkzz/storage/internal/domain/models"
	"github.com/tgkzz/storage/internal/repository/postgresql"
	"github.com/tgkzz/storage/pkg/logger"
	"log/slog"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Item struct {
	log  *slog.Logger
	repo postgresql.Repo
}

type ItemService interface {
	CreateItem(ctx context.Context, opts ...models.OptionFunc) error
	UpdateItem(ctx context.Context, itemId int64, opts ...models.OptionFunc) error
	DeleteItemById(ctx context.Context, itemId int64) error
	GetItemById(ctx context.Context, itemId int64) (models.Item, error)
}

func New(log *slog.Logger, repo postgresql.Repo) (ItemService, error) {
	return &Item{log: log, repo: repo}, nil
}

func (i *Item) CreateItem(ctx context.Context, opts ...models.OptionFunc) error {
	const op = "item.CreateItem"

	item := models.NewItem(opts...)

	l := i.log.With(
		slog.String("op", op),
		slog.String("Name", item.Name),
	)

	l.Info("create item request")

	if _, err := i.repo.CreateNewItem(ctx, *item); err != nil {
		l.Error("failed to create new item", logger.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (i *Item) UpdateItem(ctx context.Context, itemId int64, opts ...models.OptionFunc) error {
	const op = "item.UpdateItem"

	item := models.NewItem()

	models.SetId(itemId)(item)

	for _, opt := range opts {
		opt(item)
	}

	l := i.log.With(
		slog.String("op", op),
		slog.Int64("id", item.ID),
	)

	l.Info("update item request")

	r := dto.NewUpdateItem(item.ID).
		SetDesc(item.Desc).
		SetPrice(item.Price).
		SetQuantity(item.Quantity).
		SetCurrency(item.Currency).
		SetName(item.Name)

	if err := i.repo.UpdateItemById(ctx, *r); err != nil {
		l.Error("error while updating item by id", logger.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (i *Item) DeleteItemById(ctx context.Context, itemId int64) error {
	const op = "item.DeleteItemById"

	l := i.log.With(
		slog.String("op", op),
		slog.Int64("item_id", itemId),
	)

	l.Info("delete item request")

	if err := i.repo.DeleteItem(ctx, itemId); err != nil {
		l.Error("Error while deleting item", logger.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (i *Item) GetItemById(ctx context.Context, itemId int64) (models.Item, error) {
	const op = "item.GetItemById"

	l := i.log.With(
		slog.String("op", op),
		slog.Int64("item_id", itemId),
	)

	l.Info("get item by id request")

	res, err := i.repo.GetItemById(ctx, itemId)
	if err != nil {
		l.Error("error while getting item by id", logger.Err(err))
		return models.Item{}, fmt.Errorf("%s: %w", op, err)
	}

	return *res, nil
}
