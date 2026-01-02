package domain

type Product struct {
	Id       *string
	Name     *string
	Unit     *string
	Price    *float64
	Stock    *int
	Category []ProductCategory
	Audit
}

func (Product) TableName() string {
	return "product"
}
