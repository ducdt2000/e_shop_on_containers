package order_aggregate

import (
	"time"

	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
	array "github.com/ducdt2000/e_shop_on_containers/shared/helper"
)

type Order struct {
	seed_work.IAggregateRoot
	seed_work.Entity
	_orderDate       time.Time
	_address         Address
	_buyerId         int
	_orderStatus     OrderStatus
	_description     string
	_isDraft         bool
	_orderItems      []OrderItem
	_paymentMethodId int
}

func NewDraft() *Order {
	order := NewEmptyOrder()
	order._isDraft = true
	return order
}

func NewEmptyOrder() *Order {
	order := Order{_isDraft: false, _orderItems: []OrderItem{}}
	return &order
}

func NewOrder(
	userId string,
	userName string,
	address Address,
	cardTypeId int,
	cardNumber string,
	cardSecurityNumber string,
	cardHolderName string,
	cardExpiration time.Time,
	buyerId *int,
	paymentMethodId *int,
) (*Order, error) {
	entity := seed_work.NewEntity()

	order := Order{Entity: *entity}
	if buyerId != nil {
		order._buyerId = *buyerId
	}
	if paymentMethodId != nil {
		order._paymentMethodId = *paymentMethodId
	}
	order._orderStatus = OrderStatus(Submitted)
	order._orderDate = time.Now().UTC()
	order._address = address

	err := order._addOrderStartedDomainEvent(userId, userName, cardTypeId, cardNumber, cardSecurityNumber, cardHolderName, cardExpiration)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o Order) GetAddress() *Address {
	return &o._address
}

func (o Order) GetBuyerId() int {
	return o._buyerId
}

func (o Order) GetOrderStatus() OrderStatus {
	return o._orderStatus
}

func (o Order) GetOrderItems() []OrderItem {
	return o._orderItems
}

func (o *Order) SetPaymentMethodId(paymentMethodId int) {
	o._paymentMethodId = paymentMethodId
}

func (o *Order) SetBuyerId(buyerId int) {
	o._buyerId = buyerId
}

func (o *Order) AddOrderItem(
	productId int,
	productName string,
	unitPrice float32,
	discount float32,
	pictureUrl string,
	units *int,
) error {
	if units == nil {
		defaultUnit := 1
		units = &defaultUnit
	}

	existedOrderForProduct, err := array.Find(o._orderItems, func(oi OrderItem) bool {
		return oi._productId == productId
	})

	if err == nil {
		if discount > existedOrderForProduct.GetCurrentDiscount() {
			existedOrderForProduct.SetNewDiscount(discount)
		}

		existedOrderForProduct.AddUnits(*units)
	} else {
		orderItem, err := NewOrderItem(productId, productName, unitPrice, discount, pictureUrl, *units)
		if err != nil {
			return err
		}
		o._orderItems = append(o._orderItems, *orderItem)
	}

	return nil
}

func (o Order) GetTotal() float32 {
	return array.Reduce(o._orderItems, func(prev float32, curr OrderItem) float32 {
		return prev + (curr._unitPrice * float32(curr._units))
	}, 0)
}

func (o *Order) SetAwaitingValidationStatus() {
	if o._orderStatus == OrderStatus(Submitted) {
		//add domain OrderStatusChangedToAwaitingValidationDomainEvent
		o._orderStatus = OrderStatus(AwaitingValidation)
	}
}

func (o *Order) _addOrderStartedDomainEvent(userId string, userName string, cardTypeId int, cardNumber string, cardSecurityNumber string,
	cardHolderName string, cardExpiration time.Time) error {
	return nil
}
