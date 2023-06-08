package order_aggregate

type OrderStatus string

const (
	Submitted          OrderStatus = "Submitted"
	AwaitingValidation OrderStatus = "AwaitingValidation"
	StockConfirmed     OrderStatus = "StockConfirmed"
	Paid               OrderStatus = "Paid"
	Shipped            OrderStatus = "Shipped"
	Cancelled          OrderStatus = "Cancelled"
)

func List() []OrderStatus {
	return []OrderStatus{Submitted, AwaitingValidation, StockConfirmed, Paid, Shipped, Cancelled}
}
