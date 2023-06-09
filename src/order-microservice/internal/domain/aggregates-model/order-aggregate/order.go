package order_aggregate

import (
	"time"

	"github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/events"
	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
	array "github.com/ducdt2000/e_shop_on_containers/shared/helper"
)

type Order struct {
	seed_work.IAggregateRoot
	seed_work.Entity
	orderDate       time.Time
	address         Address
	buyerId         int
	orderStatus     OrderStatus
	description     string
	isDraft         bool
	orderItems      []OrderItem
	paymentMethodId int
}

func NewDraft() *Order {
	order := NewEmptyOrder()
	order.isDraft = true
	return order
}

func NewEmptyOrder() *Order {
	order := Order{isDraft: false, orderItems: []OrderItem{}}
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
		order.buyerId = *buyerId
	}
	if paymentMethodId != nil {
		order.paymentMethodId = *paymentMethodId
	}
	order.orderStatus = OrderStatus(Submitted)
	order.orderDate = time.Now().UTC()
	order.address = address

	err := order.addOrderStartedDomainEvent(userId, userName, cardTypeId, cardNumber, cardSecurityNumber, cardHolderName, cardExpiration)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o Order) GetAddress() *Address {
	return &o.address
}

func (o Order) GetBuyerId() int {
	return o.buyerId
}

func (o Order) GetOrderStatus() OrderStatus {
	return o.orderStatus
}

func (o Order) GetOrderItems() []OrderItem {
	return o.orderItems
}

func (o *Order) SetPaymentMethodId(paymentMethodId int) {
	o.paymentMethodId = paymentMethodId
}

func (o *Order) SetBuyerId(buyerId int) {
	o.buyerId = buyerId
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

	existedOrderForProduct, err := array.Find(o.orderItems, func(oi OrderItem) bool {
		return oi.productId == productId
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
		o.orderItems = append(o.orderItems, *orderItem)
	}

	return nil
}

func (o Order) GetTotal() float32 {
	return array.Reduce(o.orderItems, func(prev float32, curr OrderItem) float32 {
		return prev + (curr.unitPrice * float32(curr.units))
	}, 0)
}

func (o *Order) SetAwaitingValidationStatus() {
	if o.orderStatus == OrderStatus(Submitted) {
		o.AddDomainEvent(events.NewOrderStatusChangedToAwaitingValidationDomainEvent(o.GetId(), o.GetOrderItems()))
		o.orderStatus = OrderStatus(AwaitingValidation)
	}
}

func (o *Order) addOrderStartedDomainEvent(userId string, userName string, cardTypeId int, cardNumber string, cardSecurityNumber string,
	cardHolderName string, cardExpiration time.Time) error {
	return nil
}
