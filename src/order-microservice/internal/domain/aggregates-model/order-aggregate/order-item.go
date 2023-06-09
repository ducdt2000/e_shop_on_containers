package order_aggregate

import (
	"errors"

	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type OrderItem struct {
	seed_work.Entity
	productName string
	pictureUrl  string
	unitPrice   float32
	discount    float32
	units       int
	productId   int
}

func NewEmptyOrderItem() (*OrderItem, error) {
	return &OrderItem{}, nil
}

func NewOrderItem(
	productId int,
	productName string,
	unitPrice float32,
	discount float32,
	pictureUrl string,
	units int,
) (*OrderItem, error) {
	if units <= 0 {
		return nil, errors.New("invalid number of units")
	}

	if (unitPrice * float32(units)) < discount {
		return nil, errors.New("the total of order item is lower than applied discount")
	}

	return &OrderItem{
		productId:   productId,
		productName: productName,
		unitPrice:   unitPrice,
		pictureUrl:  pictureUrl,
		discount:    discount,
		units:       units,
	}, nil
}

func NewOrderItemWithoutUnits(
	productId int,
	productName string,
	unitPrice float32,
	discount float32,
	pictureUrl string,
) (*OrderItem, error) {
	return NewOrderItem(productId, productName, unitPrice, discount, pictureUrl, 1)
}

func (oi OrderItem) GetPictureUrl() string {
	return oi.pictureUrl
}

func (oi OrderItem) GetCurrentDiscount() float32 {
	return oi.discount
}

func (oi OrderItem) GetUnitPrice() float32 {
	return oi.unitPrice
}

func (oi OrderItem) GetItemProductName() string {
	return oi.productName
}

func (oi OrderItem) GetProductId() int {
	return oi.productId
}

func (oi OrderItem) GetUnits() int {
	return oi.units
}

func (oi *OrderItem) SetNewDiscount(discount float32) error {
	if discount < 0 {
		return errors.New("discount is not valid")
	}

	oi.discount = discount
	return nil
}

func (oi *OrderItem) AddUnits(units int) error {
	if units < 0 {
		return errors.New("units is not valid")
	}

	oi.units += units
	return nil
}
