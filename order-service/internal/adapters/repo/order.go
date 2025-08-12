package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/cripplemymind9/order-service/internal/domain/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const uniqueViolationCode = "23505"

func (q *queries) CreateOrder(ctx context.Context, order entity.Order) (int64, error) {
	const query = `
		INSERT INTO orders (
			user_id,
			product_id,
			quantity,
			total
		) VALUES (
			$1, $2, $3, $4
		) RETURNING id
	`

	var id int64
	err := q.db.QueryRow(ctx, query, order.UserID, order.ProductID, order.Quantity, order.Total).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolationCode {
			return 0, entity.ErrOrderAlreadyExists
		}
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	return id, nil
}

func (q *queries) GetOrderByID(ctx context.Context, id int64) (entity.Order, error) {
	const query = `
		SELECT
			id,
			user_id,
			product_id,
			quantity,
			total,
			status,
			created_at,
			updated_at
		FROM orders
		WHERE id = $1
	`

	row := q.db.QueryRow(ctx, query, id)

	var out entity.Order

	err := row.Scan(
		&out.ID,
		&out.UserID,
		&out.ProductID,
		&out.Quantity,
		&out.Total,
		&out.Status,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Order{}, entity.ErrOrderNotFound
		}
		return entity.Order{}, fmt.Errorf("get order by ID storage err: %w", err)
	}

	return out, nil
}
