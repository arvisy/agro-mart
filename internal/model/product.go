package model

type Product struct {
	Id       *string          `json:"id"`
	Name     *string          `json:"name" binding:"required"`
	Category *ProductCategory `json:"category" binding:"required"`
	Unit     *string          `json:"unit"`
	Price    *float64         `json:"price" binding:"required"`
	Stock    *int             `json:"stock" binding:"required"`
}
