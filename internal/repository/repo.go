package repository

import (
	"context"
	"time"

	"github.com/tgkzz/storage/internal/models"
	"github.com/tgkzz/storage/internal/repository/postgresql"
)

type Repository struct {
	ItemRepository
	CurrencyRepository
	postgresql *postgresql.PostgreRepo
}

type ItemRepository interface {
	GetItemById(ctx context.Context, id string) (*models.Item, error)
	InsertItem(ctx context.Context, item *models.Item) error
	DeleteItemById(ctx context.Context, id string) error
}

type CurrencyRepository interface {
	GetCurrencyById(ctx context.Context, id string) (*models.Currency, error)
	AddCurrency(ctx context.Context, currency models.Currency) error
}

func (r *Repository) GetItemById(ctx context.Context, id string) (*models.Item, error) {
	return r.postgresql.ItemsRepo.GetItemById(ctx, id)
}

func (r *Repository) InsertItem(ctx context.Context, item *models.Item) error {
	return r.postgresql.ItemsRepo.InsertItem(ctx, item)
}

func (r *Repository) DeleteItemById(ctx context.Context, id string) error {
	return r.postgresql.ItemsRepo.DeleteItemById(ctx, id)
}

func (r *Repository) GetCurrencyById(ctx context.Context, id string) (*models.Currency, error) {
	return r.postgresql.CurrencyRepo.GetCurrencyById(ctx, id)
}

func (r *Repository) AddCurrency(ctx context.Context, currency models.Currency) error {
	return r.postgresql.CurrencyRepo.AddCurrency(ctx, currency)
}

func NewPostgresRepository(c context.Context, url string) (*Repository, error) {
	r := postgresql.NewEmptyPostgreRepo()

	ctx, cancel := context.WithTimeout(c, 30*time.Second)
	defer cancel()

	if err := r.Connect(ctx, url); err != nil {
		return nil, err
	}

	return &Repository{
		postgresql: r,
	}, nil
}
