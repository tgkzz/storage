package items

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5"
	"github.com/tgkzz/storage/internal/models"
	"github.com/tgkzz/storage/internal/repository/errors"
)

type ItemsRepo struct {
	conn *pgx.Conn
}

func NewItemsRepo(conn *pgx.Conn) *ItemsRepo {
	return &ItemsRepo{
		conn: conn,
	}
}

const itemTable = "items"

func (ir *ItemsRepo) GetItemById(ctx context.Context, id string) (*models.Item, error) {
	sql, _, err := goqu.From(itemTable).Prepared(true).
		Where(goqu.Ex{
			"id": id,
		}).ToSQL()
	if err != nil {
		return nil, err
	}

	var item models.Item
	if err = ir.conn.QueryRow(ctx, sql, id).
		Scan(&item.Id, &item.Name, &item.Price, &item.Quantity); err != nil {
		return nil, err
	}

	return &item, nil
}

func (ir *ItemsRepo) InsertItem(ctx context.Context, item *models.Item) error {
	sql, _, err := goqu.Insert(itemTable).Prepared(true).
		Rows(goqu.Record{
			"id":          item.Id,
			"name":        item.Name,
			"quantity":    item.Quantity,
			"price":       item.Price.Price,
			"currency_id": item.Price.CurrencyId,
		}).
		Returning("id").ToSQL()
	if err != nil {
		return err
	}

	_, err = ir.conn.Exec(ctx, sql)
	return err
}

func (ir *ItemsRepo) DeleteItemById(ctx context.Context, id string) error {
	sql, _, err := goqu.Delete(itemTable).Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return err
	}

	res, err := ir.conn.Exec(ctx, sql)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return errors.ErrNotFound
	}

	return nil
}
