package product

type Product struct {
	Id    *string
	Name  *string
	Unit  *string
	Price *float64
	Stock *int
}

func (Product) TableName() string {
	return "product"
}
