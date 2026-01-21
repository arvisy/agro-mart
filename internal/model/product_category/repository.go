package productcategory

type ProductCategory struct {
	Id   *int
	Name *string
}

func (ProductCategory) TableName() string {
	return "product_category"
}
