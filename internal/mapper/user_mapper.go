package mapper

import (
	"agro-mart/internal/domain"
	"agro-mart/internal/model"
)

func UserDomainToUserModel(d *domain.User) *model.User {
	return &model.User{
		Id:   d.Id,
		Name: d.Name,
	}
}
