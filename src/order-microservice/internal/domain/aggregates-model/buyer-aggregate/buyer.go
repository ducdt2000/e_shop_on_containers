package buyer_aggregate

import (
	"errors"
	"time"

	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type Buyer struct {
	seed_work.Entity
	seed_work.IAggregateRoot
	UUID            string
	Name            string
	_paymentMethods []PaymentMethod
}

func NewBuyer() (*Buyer, error) {
	entity := seed_work.NewEntity()
	return &Buyer{_paymentMethods: []PaymentMethod{}, Entity: *entity}, nil
}

func NewBuyerWithName(uuid string, name string) (*Buyer, error) {
	if uuid == "" {
		return nil, errors.New("uuid is empty")
	}
	if name == "" {
		return nil, errors.New("name is empty")
	}
	entity := seed_work.NewEntity()
	return &Buyer{UUID: uuid, Name: name, _paymentMethods: []PaymentMethod{}, Entity: *entity}, nil

}

func (b Buyer) GetPaymentMethods() []PaymentMethod {
	return b._paymentMethods
}

func (b Buyer) VerifyOrAddPaymentMethod(
	cardType CardType,
	alias string,
	cardNumber string,
	securityNumber string,
	cardHolderName string,
	expiration time.Time,
) {
	//here
}
