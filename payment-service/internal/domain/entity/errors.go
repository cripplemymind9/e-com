package entity

import "errors"

var (
	ErrNotEnoughBalance       = errors.New("not enough available balance")
	ErrNotEnoughReservedFunds = errors.New("not enough reserved funds")
	ErrUserNotFound           = errors.New("user not found")
)
