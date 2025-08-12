package server

import (
	"context"
	"errors"

	"github.com/cripplemymind9/payment-service/internal/domain/entity"
	"github.com/cripplemymind9/payment-service/internal/domain/usecase"
	"github.com/cripplemymind9/payment-service/pkg/api/v1"
)

func (s *Server) ReserveUserBalance(
	ctx context.Context,
	req *api.ReserveUserBalanceRequest,
) (*api.ReserveUserBalanceResponse, error) {
	err := s.dependencies.reserveUserBalanceUseCase.ReserveBalance(ctx, usecase.ReserveBalanceDTO{
		UserID: req.GetUserId(),
		Amount: req.GetAmount(),
	})

	if err != nil {
		if errors.Is(err, entity.ErrNotEnoughBalance) {
			return &api.ReserveUserBalanceResponse{
				Status: api.ResponseStatus_INSUFFICIENT_QUANTITY,
			}, err
		}
		return &api.ReserveUserBalanceResponse{
			Status: api.ResponseStatus_INTERNAL_ERROR,
		}, err
	}

	return &api.ReserveUserBalanceResponse{
		Status: api.ResponseStatus_SUCCESS,
	}, nil
}

func (s *Server) CompensateUserBalance(
	ctx context.Context,
	req *api.CompensateUserBalanceRequest,
) (*api.CompensateUserBalanceResponse, error) {
	err := s.dependencies.cancelReservationUserBalanceUseCase.CancelReservation(ctx, usecase.CancelReservationDTO{
		UserID: req.GetUserId(),
		Amount: req.GetAmount(),
	})

	if err != nil {
		if errors.Is(err, entity.ErrNotEnoughReservedFunds) {
			return &api.CompensateUserBalanceResponse{
				Status: api.ResponseStatus_INSUFFICIENT_QUANTITY,
			}, err
		}
		return &api.CompensateUserBalanceResponse{
			Status: api.ResponseStatus_INTERNAL_ERROR,
		}, err
	}

	return &api.CompensateUserBalanceResponse{
		Status: api.ResponseStatus_SUCCESS,
	}, nil
}
