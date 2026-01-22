package service

import (
	"agro-mart/internal/mapper"
	"agro-mart/internal/model"
	"agro-mart/internal/repository"
)

type UserService interface {
	GetAllUser(id string) (*model.User, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (impl UserServiceImpl) GetAllUser(id string) (*model.User, error) {
	user, err := impl.UserRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	userMapping := mapper.UserDomainToUserModel(user)

	return userMapping, nil
}
