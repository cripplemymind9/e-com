package server

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cripplemymind9/order-service/internal/domain/usecase"
	"github.com/cripplemymind9/order-service/pkg/api/v1"
)

func (s *Server) CreateOrder(
	ctx context.Context, req *api.CreateOrderRequest,
) (*api.CreateOrderResponse, error) {
	dto := usecase.CreateOrderDTO{
		UserID:    req.GetUserId(),
		ProductID: req.GetProductId(),
		Quantity:  req.GetQuantity(),
		Total:     req.GetTotal(),
	}

	id, err := s.dependencies.createOrderUseCase.CreateOrder(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &api.CreateOrderResponse{OrderId: id}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *api.GetOrderRequest) (*api.GetOrderResponse, error) {
	orderDetails, err := s.dependencies.getOrderDetails.GetOrderDetails(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}

	status := mapOrderStatus(orderDetails.Status)

	return &api.GetOrderResponse{
		OrderId:   orderDetails.ID,
		UserId:    orderDetails.UserID,
		ProductId: orderDetails.ProductID,
		Quantity:  orderDetails.Quantity,
		Total:     orderDetails.Total,
		Status:    status,
		CreatedAt: timestamppb.New(orderDetails.CreatedAt),
		UpdatedAt: timestamppb.New(orderDetails.UpdatedAt),
	}, nil
}

func mapOrderStatus(status string) api.OrderStatus {
	switch status {
	case "pending":
		return api.OrderStatus_ORDER_STATUS_PENDING
	case "processing":
		return api.OrderStatus_ORDER_STATUS_PROCESSING
	case "approved":
		return api.OrderStatus_ORDER_STATUS_APPROVED
	case "rejected":
		return api.OrderStatus_ORDER_STATUS_REJECTED
	case "completed":
		return api.OrderStatus_ORDER_STATUS_COMPLETED
	default:
		return api.OrderStatus_ORDER_STATUS_UNSPECIFIED
	}
}
