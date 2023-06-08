package events

import (
	"time"

	buyer_aggregate "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/aggregates-model/buyer-aggregate"
	order_aggregate "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/aggregates-model/order-aggregate"
	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type OrderStartedDomainEvent struct {
	seed_work.INotification
	_userId             string
	_userName           string
	_cardType           buyer_aggregate.CardType
	_cardNumber         string
	_cardSecurityNumber string
	_cardHolderName     string
	_cardExpiration     time.Time
	_order              order_aggregate.Order
}

func NewOrderStartedDomainEvent(
	userId string,
	userName string, cardType buyer_aggregate.CardType, cardNumber string,
	cardSecurityNumber string, cardHolderName string, cardExpiration time.Time,
	order order_aggregate.Order,
) *OrderStartedDomainEvent {
	return &OrderStartedDomainEvent{
		_userId:             userId,
		_userName:           userName,
		_cardType:           cardType,
		_cardNumber:         cardNumber,
		_cardSecurityNumber: cardSecurityNumber,
		_cardHolderName:     cardHolderName,
		_cardExpiration:     cardExpiration,
		_order:              order,
	}
}

func (e OrderStartedDomainEvent) GetUserId() string {
	return e._userId
}

func (e OrderStartedDomainEvent) GetUserName() string {
	return e._userName
}

func (e OrderStartedDomainEvent) GetCardType() buyer_aggregate.CardType {
	return e._cardType
}

func (e OrderStartedDomainEvent) GetCardNumber() string {
	return e._cardNumber
}

func (e OrderStartedDomainEvent) GetCardSecurityNumber() string {
	return e._cardSecurityNumber
}

func (e OrderStartedDomainEvent) GetCardHolderName() string {
	return e._cardHolderName
}

func (e OrderStartedDomainEvent) GetCardExpiration() time.Time {
	return e._cardExpiration
}

func (e OrderStartedDomainEvent) GetOrder() order_aggregate.Order {
	return e._order
}
