package events

import (
	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type BuyerAndPaymentMethodVerifiedDomainEvent struct {
	seed_work.INotification
	buyer   any
	payment any
	order   any
}

func NewBuyerAndPaymentMethodVerifiedDomainEvent(
	buyer any,
	payment any,
	order any,
) *BuyerAndPaymentMethodVerifiedDomainEvent {
	return &BuyerAndPaymentMethodVerifiedDomainEvent{
		buyer:   buyer,
		payment: payment,
		order:   order,
	}
}

func (e BuyerAndPaymentMethodVerifiedDomainEvent) GetBuyer() any {
	return e.buyer
}

func (e BuyerAndPaymentMethodVerifiedDomainEvent) GetPaymentMethod() any {
	return e.payment
}

func (e BuyerAndPaymentMethodVerifiedDomainEvent) GetOrder() any {
	return e.order
}
