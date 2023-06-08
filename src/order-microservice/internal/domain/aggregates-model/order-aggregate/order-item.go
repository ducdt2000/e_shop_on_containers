package order_aggregate

import (
	"errors"

	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type OrderItem struct {
	seed_work.Entity
	_productName string
	_pictureUrl  string
	_unitPrice   float32
	_discount    float32
	_units       int
	_productId   int
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
		_productId:   productId,
		_productName: productName,
		_unitPrice:   unitPrice,
		_pictureUrl:  pictureUrl,
		_discount:    discount,
		_units:       units,
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
	return oi._pictureUrl
}

func (oi OrderItem) GetCurrentDiscount() float32 {
	return oi._discount
}

func (oi OrderItem) GetUnitPrice() float32 {
	return oi._unitPrice
}

func (oi OrderItem) GetItemProductName() string {
	return oi._productName
}

func (oi OrderItem) GetProductId() int {
	return oi._productId
}

func (oi OrderItem) GetUnits() int {
	return oi._units
}

func (oi *OrderItem) SetNewDiscount(discount float32) error {
	if discount < 0 {
		return errors.New("discount is not valid")
	}

	oi._discount = discount
	return nil
}

func (oi *OrderItem) AddUnits(units int) error {
	if units < 0 {
		return errors.New("units is not valid")
	}

	oi._units += units
	return nil
}
