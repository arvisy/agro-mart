package domain

import "time"

type Product struct {
	Id         *string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       *string
	CategoryId *int
	Unit       *string
	Price      *float64
	Stock      *int
	CreatedAt  *time.Time
}

func (Product) TableName() string {
	return "product"
}
