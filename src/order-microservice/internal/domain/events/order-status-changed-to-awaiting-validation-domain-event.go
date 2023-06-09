package events

type OrderStatusChangedToAwaitingValidationDomainEvent struct {
	orderId    int
	orderItems []any
}

func NewOrderStatusChangedToAwaitingValidationDomainEvent(orderId int, orderItems []any) *OrderStatusChangedToAwaitingValidationDomainEvent {
	return &OrderStatusChangedToAwaitingValidationDomainEvent{
		orderId:    orderId,
		orderItems: orderItems,
	}
}

func (e OrderStatusChangedToAwaitingValidationDomainEvent) GetOrderId() int {
	return e.orderId
}

func (e OrderStatusChangedToAwaitingValidationDomainEvent) GetOrderItems() []any {
	return e.orderItems
}
