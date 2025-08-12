package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/cripplemymind9/payment-service/internal/domain/entity"
	"github.com/jackc/pgx/v5"
)

func (q *queries) GetUserBalanceByID(ctx context.Context, id int64) (entity.UserBalance, error) {
	const query = `
		SELECT
			user_id,
			total_balance,
			reserved_balance,
			available_balance
		FROM user_balances
		WHERE user_id = $1
	`

	row := q.db.QueryRow(ctx, query, id)

	var out entity.UserBalance

	err := row.Scan(
		&out.UserID,
		&out.TotalBalance,
		&out.ReservedBalance,
		&out.AvailableBalance,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.UserBalance{}, entity.ErrUserNotFound
		}
		return entity.UserBalance{}, fmt.Errorf("get user balance by userID storage err: %w", err)
	}

	return out, nil
}

func (q *queries) ReserveBalance(ctx context.Context, userID, amount int64) error {
	const query = `
		UPDATE user_balances
		SET
			reserved_balance = reserved_balance + $1,
			available_balance = available_balance - $1
		WHERE
			user_id = $2
			AND available_balance >= $1
	`

	tag, err := q.db.Exec(ctx, query, amount, userID)
	if err != nil {
		return fmt.Errorf("reserve user_balance storage err: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return entity.ErrNotEnoughBalance
	}

	return nil
}

func (q *queries) CancelReservation(ctx context.Context, userID, amount int64) error {
	const query = `
		UPDATE user_balances
		SET
			reserved_balance = reserved_balance - $1,
			available_balance = available_balance + $1
		WHERE
			user_id = $2
			AND reserved_balance >= $1
	`

	tag, err := q.db.Exec(ctx, query, amount, userID)
	if err != nil {
		return fmt.Errorf("cancel reservation storage err: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return entity.ErrNotEnoughReservedFunds
	}

	return nil
}
