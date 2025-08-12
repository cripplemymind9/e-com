package queu

const (
	EventTypeOrderCreated = "order_created"
)

type OrderItem struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type OrderCreatedEvent struct {
	EventType string      `json:"event_type"`
	OrderID   int64       `json:"order_id"`
	UserID    int64       `json:"user_id"`
	Items     []OrderItem `json:"items"`
	Total     float64     `json:"total"`
}
