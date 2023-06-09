package events

type OrderStatusChangedToStockConfirmedDomainEvent struct {
	orderId int
}

func NewOrderStatusChangedToStockConfirmedDomainEvent(orderId int, orderItems []any) *OrderStatusChangedToStockConfirmedDomainEvent {
	return &OrderStatusChangedToStockConfirmedDomainEvent{
		orderId: orderId,
	}
}

func (e OrderStatusChangedToStockConfirmedDomainEvent) GetOrderId() int {
	return e.orderId
}
