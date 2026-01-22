package model

type Product struct {
	Id         *string  `json:"id"`
	Name       *string  `json:"name" binding:"required"`
	CategoryId *int     `json:"categoryId" binding:"required"`
	Unit       *string  `json:"unit"`
	Price      *float64 `json:"price" binding:"required"`
	Stock      *int     `json:"stock" binding:"required"`
}
