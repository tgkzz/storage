package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/tgkzz/storage/internal/repository/errors"
	"github.com/tgkzz/storage/internal/repository/postgresql/currency"
	"github.com/tgkzz/storage/internal/repository/postgresql/items"
)

type PostgreRepo struct {
	CurrencyRepo *currency.CurrencyRepo
	ItemsRepo    *items.ItemsRepo
	conn         *pgx.Conn
	connected    bool
}

func NewEmptyPostgreRepo() *PostgreRepo {
	return &PostgreRepo{
		connected: false,
	}
}

func (p *PostgreRepo) Connect(ctx context.Context, url string) error {
	connect, err := pgx.Connect(ctx, url)
	if err != nil {
		return err
	}
	if err = connect.Ping(ctx); err != nil {
		return err
	}

	p.CurrencyRepo = currency.NewCurrencyRepo(connect)
	p.ItemsRepo = items.NewItemsRepo(connect)
	p.conn = connect
	p.connected = true

	return nil
}

func (p *PostgreRepo) Disconnect(ctx context.Context) error {
	if p.CurrencyRepo == nil || p.ItemsRepo == nil {
		return errors.ErrNotConnected
	}

	return p.conn.Close(ctx)
}
