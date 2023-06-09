package order_aggregate

import seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"

type Address struct {
	seed_work.ValueObject
	street  string
	city    string
	state   string
	country string
	zipCode string
}

func NewEmptyAddress() *Address {
	return &Address{}
}

func NewAddress(street string, city string, state string, country string, zipCode string) *Address {
	return &Address{
		street:  street,
		city:    city,
		state:   state,
		country: country,
		zipCode: zipCode,
	}
}

func (addr Address) GetStreet() string {
	return addr.street
}

func (addr Address) GetCity() string {
	return addr.city
}

func (addr Address) GetState() string {
	return addr.state
}

func (addr Address) GetCountry() string {
	return addr.country
}

func (addr Address) GetZipCode() string {
	return addr.zipCode
}
