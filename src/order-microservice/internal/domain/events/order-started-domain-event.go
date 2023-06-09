package events

import (
	"time"

	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type OrderStartedDomainEvent struct {
	seed_work.INotification
	userId             string
	userName           string
	cardType           any
	cardNumber         string
	cardSecurityNumber string
	cardHolderName     string
	cardExpiration     time.Time
	order              any
}

func NewOrderStartedDomainEvent(
	userId string,
	userName string, cardType any, cardNumber string,
	cardSecurityNumber string, cardHolderName string, cardExpiration time.Time,
	order any,
) *OrderStartedDomainEvent {
	return &OrderStartedDomainEvent{
		userId:             userId,
		userName:           userName,
		cardType:           cardType,
		cardNumber:         cardNumber,
		cardSecurityNumber: cardSecurityNumber,
		cardHolderName:     cardHolderName,
		cardExpiration:     cardExpiration,
		order:              order,
	}
}

func (e OrderStartedDomainEvent) GetUserId() string {
	return e.userId
}

func (e OrderStartedDomainEvent) GetUserName() string {
	return e.userName
}

func (e OrderStartedDomainEvent) GetCardType() any {
	return e.cardType
}

func (e OrderStartedDomainEvent) GetCardNumber() string {
	return e.cardNumber
}

func (e OrderStartedDomainEvent) GetCardSecurityNumber() string {
	return e.cardSecurityNumber
}

func (e OrderStartedDomainEvent) GetCardHolderName() string {
	return e.cardHolderName
}

func (e OrderStartedDomainEvent) GetCardExpiration() time.Time {
	return e.cardExpiration
}

func (e OrderStartedDomainEvent) GetOrder() any {
	return e.order
}
