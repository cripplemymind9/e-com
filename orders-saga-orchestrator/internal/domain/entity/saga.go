package entity

import "time"

type SagaStatus string

const (
	SagaStatusUnspecified SagaStatus = "unspecified"
	SagaStatusPending     SagaStatus = "pending"
	SagaStatusCompleted   SagaStatus = "completed"
	SagaStatusFailed      SagaStatus = "failed"
	SagaStatusCompensated SagaStatus = "compensated"
)

type Saga struct {
	ID        int64      `json:"id"`
	OrderID   int64      `json:"order_id"`
	Status    SagaStatus `json:"status"`
	Steps     []Step     `json:"steps"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Step struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
