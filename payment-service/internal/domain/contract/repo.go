package contract

import (
	"context"

	"github.com/cripplemymind9/payment-service/internal/domain/entity"
)

type RepoTransactor interface {
	InTx(ctx context.Context, f func(tx TxRepo) error) error
}

type TxRepo interface {
	GetUserBalanceByID(ctx context.Context, id int64) (entity.UserBalance, error)
	ReserveBalance(ctx context.Context, userID, amount int64) error
	CancelReservation(ctx context.Context, userID, amount int64) error
}
