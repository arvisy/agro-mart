package repository

import (
	"agro-mart/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() ([]domain.Product, error)
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func (impl ProductRepositoryImpl) GetAllProduct() ([]domain.Product, error) {
	q := impl.DB.Debug().Model(&domain.Product{})

	products := []domain.Product{}
	tx := q.Find(&products)
	if tx.Error != nil {
		return []domain.Product{}, tx.Error
	}

	return products, nil
}
