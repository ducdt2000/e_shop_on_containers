package events

import (
	buyer_aggregate "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/aggregates-model/buyer-aggregate"
	order_aggregate "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/aggregates-model/order-aggregate"
	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type BuyerAndPaymentMethodVerifiedDomainEvent struct {
	seed_work.INotification
	_buyer   buyer_aggregate.Buyer
	_payment buyer_aggregate.PaymentMethod
	_order   order_aggregate.Order
}

func NewBuyerAndPaymentMethodVerifiedDomainEvent(
	buyer buyer_aggregate.Buyer,
	payment buyer_aggregate.PaymentMethod,
	order order_aggregate.Order,
) *BuyerAndPaymentMethodVerifiedDomainEvent {
	return &BuyerAndPaymentMethodVerifiedDomainEvent{
		_buyer:   buyer,
		_payment: payment,
		_order:   order,
	}
}

func (e BuyerAndPaymentMethodVerifiedDomainEvent) GetBuyer() buyer_aggregate.Buyer {
	return e._buyer
}

func (e BuyerAndPaymentMethodVerifiedDomainEvent) GetPaymentMethod() buyer_aggregate.PaymentMethod {
	return e._payment
}

func (e BuyerAndPaymentMethodVerifiedDomainEvent) GetOrder() order_aggregate.Order {
	return e._order
}
