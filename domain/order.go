package domain

type Order struct {
	Id          *string
	UserId      *User
	Status      *string
	TotalAmount *int
	Audit
}

func (Order) TableName() string {
	return "order"
}
