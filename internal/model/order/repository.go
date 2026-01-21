package order

type Order struct {
	Id          *string
	UserId      *string
	Status      *string
	TotalAmount *int
}

func (Order) TableName() string {
	return "order"
}
