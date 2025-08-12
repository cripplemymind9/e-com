package usecase

import (
	"context"

	"github.com/cripplemymind9/payment-service/internal/domain/contract"
	"github.com/cripplemymind9/payment-service/internal/domain/entity"
)

type ReserveBalanceUseCase struct {
	transactor contract.RepoTransactor
}

func NewReserveBalanceUseCase(
	transactor contract.RepoTransactor,
) *ReserveBalanceUseCase {
	return &ReserveBalanceUseCase{
		transactor: transactor,
	}
}

func (rb *ReserveBalanceUseCase) ReserveBalance(
	ctx context.Context,
	dto ReserveBalanceDTO,
) error {
	return rb.transactor.InTx(ctx, func(tx contract.TxRepo) error {
		userBalance, err := tx.GetUserBalanceByID(ctx, dto.UserID)
		if err != nil {
			return err
		}

		if userBalance.AvailableBalance < dto.Amount {
			return entity.ErrNotEnoughBalance
		}

		return tx.ReserveBalance(ctx, dto.UserID, dto.Amount)
	})
}
