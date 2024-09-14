package item

import (
	"context"
	"github.com/tgkzz/storage/internal/domain/models"
	"log/slog"
)

type Item struct {
	log *slog.Logger
}

type ItemService interface {
	CreateItem(ctx context.Context, opts ...models.OptionFunc) error
	UpdateItem(ctx context.Context, itemId int64, opts ...models.OptionFunc) error
	DeleteItemById(ctx context.Context, itemId int64) error
	GetItemById(ctx context.Context, itemId int64) (models.Item, error)
}

func New() (*ItemService, error) {
	return nil, nil
}

func (i *Item) CreateItem(ctx context.Context, opts ...models.OptionFunc) error {
	const op = "item.CreateItem"

	item := models.NewItem(opts...)

	_ = item

	return nil
}

func (i *Item) UpdateItem(ctx context.Context, itemId int64, opts ...models.OptionFunc) error {
	const op = "item.UpdateItem"

	item := models.NewItem()

	models.SetId(itemId)(item)

	for _, opt := range opts {
		opt(item)
	}

	return nil
}

func (i *Item) DeleteItemById(ctx context.Context, itemId int64) error {
	const op = "item.DeleteItemById"

	return nil
}

func (i *Item) GetItemById(ctx context.Context, itemId int64) (models.Item, error) {
	const op = "item.GetItemById"

	return models.Item{}, nil
}
