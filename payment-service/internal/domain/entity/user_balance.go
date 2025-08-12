package entity

type UserBalance struct {
	UserID           int64
	TotalBalance     int64
	ReservedBalance  int64
	AvailableBalance int64
}
