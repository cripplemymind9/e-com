package usecase

import (
	"time"

	"github.com/cripplemymind9/orders-saga-orchestrator/internal/domain/entity"
)

type SagaDTO struct {
	ID        int64
	OrderID   int64
	Status    string
	Steps     []StepDTO
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type StepDTO struct {
	Name   string
	Status string
}

func toStepDTOs(steps []entity.Step) []StepDTO {
	dtos := make([]StepDTO, len(steps))
	for i, step := range steps {
		dtos[i] = StepDTO{
			Name:   step.Name,
			Status: step.Status,
		}
	}
	return dtos
}
