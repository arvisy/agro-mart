package mapper

import (
	"agro-mart/internal/domain"
	"agro-mart/internal/model"
)

func UserDomainToUserModel(user *domain.User) *model.User {
	return &model.User{
		Id:   user.Id,
		Name: user.Name,
	}
}

func UserModelToUserDomain(user *model.User) *domain.User {
	return &domain.User{
		Id:   user.Id,
		Name: user.Name,
		Role: user.Role,
	}
}
