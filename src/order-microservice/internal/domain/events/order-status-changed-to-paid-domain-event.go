package events

type OrderStatusChangedToPaidDomainEvent struct {
	orderId    int
	orderItems []any
}

func NewOrderStatusChangedToPaidDomainEvent(orderId int, orderItems []any) *OrderStatusChangedToPaidDomainEvent {
	return &OrderStatusChangedToPaidDomainEvent{
		orderId:    orderId,
		orderItems: orderItems,
	}
}

func (e OrderStatusChangedToPaidDomainEvent) GetOrderId() int {
	return e.orderId
}

func (e OrderStatusChangedToPaidDomainEvent) GetOrderItems() []any {
	return e.orderItems
}
