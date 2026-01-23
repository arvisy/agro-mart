package repository

import (
	"agro-mart/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	GetAllProduct() ([]domain.Product, error)
	UpsertProduct(*domain.Product) (*domain.Product, error)
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func (impl ProductRepositoryImpl) mainQuery() *gorm.DB {
	return impl.DB.Debug().Model(&domain.Product{})
}

func (impl ProductRepositoryImpl) GetAllProduct() ([]domain.Product, error) {
	products := []domain.Product{}
	tx := impl.mainQuery().Preload("ProductCategory").Find(&products)
	if tx.Error != nil {
		return []domain.Product{}, tx.Error
	}

	return products, nil
}

func (impl ProductRepositoryImpl) UpsertProduct(product *domain.Product) (*domain.Product, error) {
	tx := impl.mainQuery().Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"category_id",
			"unit",
			"price",
			"stock",
		}),
	})

	if tx.Error != nil {
		return nil, tx.Error
	}

	return product, nil
}
