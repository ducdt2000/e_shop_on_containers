package order_aggregate

import seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"

type Address struct {
	seed_work.ValueObject
	_street  string
	_city    string
	_state   string
	_country string
	_zipCode string
}

func NewEmptyAddress() *Address {
	return &Address{}
}

func NewAddress(street string, city string, state string, country string, zipCode string) *Address {
	return &Address{
		_street:  street,
		_city:    city,
		_state:   state,
		_country: country,
		_zipCode: zipCode,
	}
}

func (addr Address) GetStreet() string {
	return addr._street
}

func (addr Address) GetCity() string {
	return addr._city
}

func (addr Address) GetState() string {
	return addr._state
}

func (addr Address) GetCountry() string {
	return addr._country
}

func (addr Address) GetZipCode() string {
	return addr._zipCode
}
