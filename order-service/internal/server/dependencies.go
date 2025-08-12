package server

import "github.com/cripplemymind9/order-service/internal/domain/usecase"

type Dependencies struct {
	createOrderUseCase *usecase.CreateOrderUseCase
	getOrderDetails    *usecase.GetOrderDetailsUseCase
}

func NewDependencies(
	createOrderUseCase *usecase.CreateOrderUseCase,
	getOrderDetailsUseCase *usecase.GetOrderDetailsUseCase,
) *Dependencies {
	return &Dependencies{
		createOrderUseCase: createOrderUseCase,
		getOrderDetails:    getOrderDetailsUseCase,
	}
}
