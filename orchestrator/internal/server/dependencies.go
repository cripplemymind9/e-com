package server

import "github.com/cripplemymind9/orders-saga-orchestrator/internal/domain/usecase"

type Dependencies struct {
	getSagaByOrderIDUseCase *usecase.GetSagaByOrderIDUseCase
}

func NewDependencies(
	getSagaByOrderIDUseCase *usecase.GetSagaByOrderIDUseCase,
) *Dependencies {
	return &Dependencies{
		getSagaByOrderIDUseCase: getSagaByOrderIDUseCase,
	}
}
