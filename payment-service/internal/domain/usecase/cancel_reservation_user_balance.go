package usecase

import (
	"context"

	"github.com/cripplemymind9/payment-service/internal/domain/contract"
	"github.com/cripplemymind9/payment-service/internal/domain/entity"
)

type CancelReservationUserBalanceUseCase struct {
	transactor contract.RepoTransactor
}

func NewCancelReservationUserBalanceUseCase(
	transactor contract.RepoTransactor,
) *CancelReservationUserBalanceUseCase {
	return &CancelReservationUserBalanceUseCase{
		transactor: transactor,
	}
}

func (crub *CancelReservationUserBalanceUseCase) CancelReservation(
	ctx context.Context,
	dto CancelReservationDTO,
) error {
	return crub.transactor.InTx(ctx, func(tx contract.TxRepo) error {
		userBalance, err := tx.GetUserBalanceByID(ctx, dto.UserID)
		if err != nil {
			return err
		}

		if userBalance.ReservedBalance < dto.Amount {
			return entity.ErrNotEnoughReservedFunds
		}

		return tx.CancelReservation(ctx, dto.UserID, dto.Amount)
	})
}
