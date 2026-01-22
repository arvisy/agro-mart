package service

import (
	"agro-mart/internal/mapper"
	"agro-mart/internal/model"
	"agro-mart/internal/repository"
)

type ProductService interface {
	GetAllProduct() ([]model.Product, error)
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func (impl ProductServiceImpl) GetAllProduct() ([]model.Product, error) {
	product, err := impl.ProductRepository.GetAllProduct()
	if err != nil {
		return []model.Product{}, err
	}

	productModel := mapper.ListProductDomainToListProductModel(product)

	return productModel, nil
}
