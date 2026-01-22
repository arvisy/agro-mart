package mapper

import (
	"agro-mart/internal/domain"
	"agro-mart/internal/model"
)

func ListProductDomainToListProductModel(d []domain.Product) []model.Product {
	products := []model.Product{}

	for _, v := range d {
		products = append(products, model.Product{
			Id:         v.Id,
			Name:       v.Name,
			CategoryId: v.CategoryId,
			Unit:       v.Unit,
			Price:      v.Price,
			Stock:      v.Stock,
		})
	}

	return products
}
