package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/cripplemymind9/order-service/internal/domain/entity"
)

type coCreateOrderRepo interface {
	CreateOrder(ctx context.Context, order entity.Order) (int64, error)
}

type coOrderProducer interface {
	SendOrderCreatedEvent(ctx context.Context, orderID int64, dto CreateOrderDTO) error
}

type CreateOrderUseCase struct {
	createOrderRepo coCreateOrderRepo
	orderProducer   coOrderProducer
}

func NewCreateOrderUseCase(
	createOrderRepo coCreateOrderRepo,
	orderProducer coOrderProducer,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		createOrderRepo: createOrderRepo,
		orderProducer:   orderProducer,
	}
}

func (co *CreateOrderUseCase) CreateOrder(ctx context.Context, dto CreateOrderDTO) (int64, error) {
	order := entity.Order{
		UserID:    dto.UserID,
		ProductID: dto.ProductID,
		Quantity:  dto.Quantity,
		Total:     dto.Total,
	}

	id, err := co.createOrderRepo.CreateOrder(ctx, order)
	if err != nil && !errors.Is(err, entity.ErrOrderAlreadyExists) {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	if errEvent := co.orderProducer.SendOrderCreatedEvent(ctx, id, dto); errEvent != nil {
		return id, fmt.Errorf("failed to send order created event: %w", errEvent)
	}

	return id, nil
}
