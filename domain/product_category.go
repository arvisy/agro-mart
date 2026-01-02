package domain

type ProductCategory struct {
	Id   *int
	Name *string
	Audit
}

func (ProductCategory) TableName() string {
	return "product_category"
}
