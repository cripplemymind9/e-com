package queu

import (
	"context"
	"fmt"

	"github.com/cripplemymind9/order-service/internal/config"
	"github.com/cripplemymind9/order-service/internal/domain/usecase"
	"github.com/cripplemymind9/order-service/internal/pkg/kafka"
)

type OrderProducer struct {
	producer *kafka.Producer[OrderCreatedEvent]
}

func NewOrderProducer(cfg config.Kafka) (*OrderProducer, error) {
	producer, err := kafka.NewProducer[OrderCreatedEvent](kafka.Config{
		Hosts: cfg.OrderSaga.Hosts,
	},
		cfg.OrderSaga.Topic,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create order producer: %w", err)
	}

	return &OrderProducer{
		producer: producer,
	}, nil
}

func (op *OrderProducer) SendOrderCreatedEvent(ctx context.Context, orderID int64, dto usecase.CreateOrderDTO) error {
	msg := OrderCreatedEvent{
		EventType: EventTypeOrderCreated,
		OrderID:   orderID,
		UserID:    dto.UserID,
		Items: []OrderItem{
			{
				ProductID: dto.ProductID,
				Quantity:  int64(dto.Quantity),
			},
		},
		Total: float64(dto.Total),
	}

	if err := op.producer.Produce(ctx, msg); err != nil {
		return fmt.Errorf("failed to send order created event: %w", err)
	}

	return nil
}
