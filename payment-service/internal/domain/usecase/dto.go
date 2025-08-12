package usecase

type ReserveBalanceDTO struct {
	UserID int64
	Amount int64
}

type CancelReservationDTO struct {
	UserID int64
	Amount int64
}
