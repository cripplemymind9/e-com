package server

import (
	"context"

	"github.com/cripplemymind9/orders-saga-orchestrator/internal/domain/usecase"
	api "github.com/cripplemymind9/orders-saga-orchestrator/pkg/api/v1"
)

func (s *Server) GetOrderSaga(
	ctx context.Context,
	req *api.GetOrderSagaRequest,
) (*api.GetOrderSagaResponse, error) {
	saga, err := s.dependencies.getSagaByOrderIDUseCase.GetSagaByOrderID(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}

	return &api.GetOrderSagaResponse{
		Saga: &api.OrderSaga{
			Id:      saga.ID,
			OrderId: saga.OrderID,
			Status:  convertStatus(saga.Status),
			Steps:   convertSteps(saga.Steps),
		},
	}, nil
}

func convertStatus(status string) api.SagaStatus {
	switch status {
	case "pending":
		return api.SagaStatus_SAGA_STATUS_PENDING
	case "completed":
		return api.SagaStatus_SAGA_STATUS_COMPLETED
	case "failed":
		return api.SagaStatus_SAGA_STATUS_FAILED
	case "compensated":
		return api.SagaStatus_SAGA_STATUS_COMPENSATED
	default:
		return api.SagaStatus_SAGA_STATUS_UNSPECIFIED
	}
}

func convertSteps(steps []usecase.StepDTO) []*api.SagaStep {
	result := make([]*api.SagaStep, len(steps))
	for i, step := range steps {
		result[i] = &api.SagaStep{
			Name:   step.Name,
			Status: convertStepStatus(step.Status),
		}
	}
	return result
}

func convertStepStatus(status string) api.SagaStepStatus {
	switch status {
	case "completed":
		return api.SagaStepStatus_SAGA_STEP_STATUS_COMPLETED
	case "failed":
		return api.SagaStepStatus_SAGA_STEP_STATUS_FAILED
	default:
		return api.SagaStepStatus_SAGA_STEP_STATUS_UNSPECIFIED
	}
}
