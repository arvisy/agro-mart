package mapper

import (
	"agro-mart/internal/domain"
	"agro-mart/internal/model"
)

func ListProductDomainToListProductModel(d []domain.Product) []model.Product {
	products := []model.Product{}

	for _, v := range d {
		products = append(products, model.Product{
			Id:    v.Id,
			Name:  v.Name,
			Unit:  v.Unit,
			Price: v.Price,
			Stock: v.Stock,
			Category: &model.ProductCategory{
				Id:   v.ProductCategory.Id,
				Name: v.ProductCategory.Name,
			},
		})
	}

	return products
}
