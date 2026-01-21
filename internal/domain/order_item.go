package domain

type OrderItem struct {
	Id        *string
	OrderId   *string
	ProductId *string
	Quantity  *int
	Price     *float64
}

func (OrderItem) TableName() string {
	return "order_item"
}
