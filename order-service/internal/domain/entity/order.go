package entity

import "time"

type Order struct {
	ID        int64
	UserID    int64
	ProductID int64
	Quantity  int32
	Total     int64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
