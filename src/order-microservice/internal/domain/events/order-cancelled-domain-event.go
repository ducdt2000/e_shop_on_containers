package events

import (
	order_aggregate "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/aggregates-model/order-aggregate"
	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type OrderCancelledDomainEvent struct {
	seed_work.INotification
	_order order_aggregate.Order
}

func NewOrderCancelledDomainEvent(order order_aggregate.Order) *OrderCancelledDomainEvent {
	return &OrderCancelledDomainEvent{_order: order}
}

func (e OrderCancelledDomainEvent) GetOrder() order_aggregate.Order {
	return e._order
}
