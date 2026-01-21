package mapper

import (
	"agro-mart/internal/domain"
	"agro-mart/internal/model"
)

func ListUserDomainToListUserModel(d []domain.User) []model.User {
	res := []model.User{}

	for _, v := range d {
		res = append(res, model.User{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return res
}
