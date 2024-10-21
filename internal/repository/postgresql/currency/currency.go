package currency

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5"
	"github.com/tgkzz/storage/internal/models"
)

type CurrencyRepo struct {
	conn *pgx.Conn
}

const currencyTable = "currency"

func NewCurrencyRepo(conn *pgx.Conn) *CurrencyRepo {
	return &CurrencyRepo{
		conn: conn,
	}
}

func (cr *CurrencyRepo) GetCurrencyById(ctx context.Context, id string) (*models.Currency, error) {
	sql, _, err := goqu.From(currencyTable).Prepared(true).
		Where(goqu.Ex{
			"id": id,
		}).ToSQL()
	if err != nil {
		return nil, err
	}

	var currency models.Currency
	if err = cr.conn.QueryRow(ctx, sql).
		Scan(&currency.Id, &currency.Code, &currency.Name); err != nil {
		return nil, err
	}

	return &currency, nil
}

func (cr *CurrencyRepo) AddCurrency(ctx context.Context, currency models.Currency) error {
	sql, _, err := goqu.Insert(currencyTable).Prepared(true).Rows(
		goqu.Record{
			"id":   currency.Id,
			"code": currency.Code,
			"name": currency.Name,
		}).ToSQL()
	if err != nil {
		return err
	}

	if _, err = cr.conn.Exec(ctx, sql); err != nil {
		return err
	}

	return nil
}
