package server

import "github.com/cripplemymind9/payment-service/internal/domain/usecase"

type Dependencies struct {
	reserveUserBalanceUseCase           *usecase.ReserveBalanceUseCase
	cancelReservationUserBalanceUseCase *usecase.CancelReservationUserBalanceUseCase
}

func NewDependencies(
	reserveUserBalanceUseCase *usecase.ReserveBalanceUseCase,
	cancelReservationUserBalanceUseCase *usecase.CancelReservationUserBalanceUseCase,
) *Dependencies {
	return &Dependencies{
		reserveUserBalanceUseCase:           reserveUserBalanceUseCase,
		cancelReservationUserBalanceUseCase: cancelReservationUserBalanceUseCase,
	}
}
