package usecase

import (
	"context"
	"fmt"

	"github.com/cripplemymind9/order-service/internal/domain/entity"
)

type godOrderRepository interface {
	GetOrderByID(ctx context.Context, id int64) (entity.Order, error)
}

type GetOrderDetailsUseCase struct {
	orderRepository godOrderRepository
}

func NewGetOrderDetailsUseCase(
	orderRepository godOrderRepository,
) *GetOrderDetailsUseCase {
	return &GetOrderDetailsUseCase{
		orderRepository: orderRepository,
	}
}

func (god *GetOrderDetailsUseCase) GetOrderDetails(
	ctx context.Context, id int64,
) (OrderDetailsDTO, error) {
	orderEntity, err := god.orderRepository.GetOrderByID(ctx, id)
	if err != nil {
		return OrderDetailsDTO{}, fmt.Errorf("get order details getting order: %w", err)
	}

	return OrderDetailsDTO{
		ID:        orderEntity.ID,
		UserID:    orderEntity.UserID,
		ProductID: orderEntity.ProductID,
		Quantity:  orderEntity.Quantity,
		Total:     orderEntity.Total,
		Status:    orderEntity.Status,
		CreatedAt: orderEntity.CreatedAt,
		UpdatedAt: orderEntity.UpdatedAt,
	}, nil
}
