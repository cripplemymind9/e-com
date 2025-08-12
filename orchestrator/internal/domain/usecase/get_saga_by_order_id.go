package usecase

import (
	"context"

	"github.com/cripplemymind9/orders-saga-orchestrator/internal/domain/entity"
)

type gsSagaRepository interface {
	GetSagaByOrderID(ctx context.Context, id int64) (entity.Saga, error)
}

type GetSagaByOrderIDUseCase struct {
	sagaRepository gsSagaRepository
}

func NewGetSagaByOrderIDUseCase(
	sagaRepository gsSagaRepository,
) *GetSagaByOrderIDUseCase {
	return &GetSagaByOrderIDUseCase{
		sagaRepository: sagaRepository,
	}
}

func (gs *GetSagaByOrderIDUseCase) GetSagaByOrderID(ctx context.Context, orderID int64) (SagaDTO, error) {
	saga, err := gs.sagaRepository.GetSagaByOrderID(ctx, orderID)
	if err != nil {
		return SagaDTO{}, err
	}

	return SagaDTO{
		ID:        saga.ID,
		OrderID:   saga.OrderID,
		Status:    string(saga.Status),
		Steps:     toStepDTOs(saga.Steps),
		CreatedAt: saga.CreatedAt,
		UpdatedAt: saga.UpdatedAt,
	}, nil
}
