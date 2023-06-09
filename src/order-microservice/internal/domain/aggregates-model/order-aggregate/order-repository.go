package order_aggregate

import seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"

type IOrderRepository interface {
	seed_work.IRepository
	Add(order Order) Order
	Update(orderId int, order Order) Order
	GetOne(orderId int) Order
}
