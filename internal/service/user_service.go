package service

import (
	"agro-mart/internal/mapper"
	"agro-mart/internal/model"
	"agro-mart/internal/repository"
)

type UserService interface {
	GetAllUser() ([]model.User, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (impl UserServiceImpl) GetAllUser() ([]model.User, error) {
	users, err := impl.UserRepository.GetAllUser()
	if err != nil {
		return []model.User{}, err
	}

	userMapping := mapper.ListUserDomainToListUserModel(users)

	return userMapping, nil
}
