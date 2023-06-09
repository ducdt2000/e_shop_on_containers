package events

import (
	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type OrderShippedDomainEvent struct {
	seed_work.INotification
	order any
}

func NewOrderShippedDomainEvent(order any) *OrderCancelledDomainEvent {
	return &OrderCancelledDomainEvent{order: order}
}

func (e OrderShippedDomainEvent) GetOrder() any {
	return e.order
}
