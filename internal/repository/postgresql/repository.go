package postgresql

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tgkzz/storage/internal/domain/dto"
	"github.com/tgkzz/storage/internal/domain/models"
)

var (
	ErrNotFound = errors.New("not found")
)

type Repository struct {
	dbConn *pgxpool.Pool
}

type Repo interface {
	CreateNewItem(ctx context.Context, item models.Item) (itemId int64, err error)
	DeleteItem(ctx context.Context, itemId int64) error
	GetItemById(ctx context.Context, itemId int64) (*models.Item, error)
	UpdateItemById(ctx context.Context, item dto.UpdateItem) error
}

func (r *Repository) Close() {
	r.dbConn.Close()
}

func New(ctx context.Context, dbUrl string) (*Repository, error) {
	c, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return nil, err
	}

	if err = c.Ping(ctx); err != nil {
		return nil, err
	}

	return &Repository{dbConn: c}, nil
}

func (r *Repository) CreateNewItem(ctx context.Context, item models.Item) (itemId int64, err error) {
	const op = "repo.postgresql.CreateNewItem"

	query, args, err := sq.Insert("items").
		Columns("name", "description", "quantity", "price", "currency").
		Values(item.Name, item.Desc, item.Quantity, item.Price, item.Currency).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	if err = r.dbConn.QueryRow(ctx, query, args...).Scan(&itemId); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return
}

func (r *Repository) DeleteItem(ctx context.Context, itemId int64) error {
	const op = "repo.postgresql.DeleteItem"

	query, args, err := sq.Delete("items").
		Where(sq.Eq{"id": itemId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res, err := r.dbConn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return nil
}

func (r *Repository) GetItemById(ctx context.Context, itemId int64) (*models.Item, error) {
	const op = "repo.postgresql.GetItemById"

	query, args, err := sq.Select("*").
		From("items").
		Where(sq.Eq{"id": itemId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var res models.Item
	if err = r.dbConn.QueryRow(ctx, query, args...).Scan(&res.ID, &res.Name, &res.Desc, &res.Price, &res.Quantity); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &res, nil
}

func (r *Repository) UpdateItemById(ctx context.Context, item dto.UpdateItem) error {
	const op = "repo.postgresql.UpdateItemById"

	query, args, err := r.updateBuilder(item).ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res, err := r.dbConn.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return nil
}

func (r *Repository) updateBuilder(item dto.UpdateItem) sq.UpdateBuilder {
	queryBuilder := sq.Update("items").
		Where(sq.Eq{"id": item.Id}).
		PlaceholderFormat(sq.Dollar)

	if item.HaveName() {
		queryBuilder = queryBuilder.Set("name", *item.Name)
	}

	if item.HaveDesc() {
		queryBuilder = queryBuilder.Set("description", *item.Desc)
	}

	if item.HavePrice() {
		queryBuilder = queryBuilder.Set("price", *item.Price)
	}

	if item.HaveQuantity() {
		queryBuilder = queryBuilder.Set("quantity", *item.Quantity)
	}

	if item.HaveCurrency() {
		queryBuilder = queryBuilder.Set("currency", *item.Currency)
	}

	return queryBuilder
}
