package service

import (
	"agro-mart/internal/mapper"
	"agro-mart/internal/model"
	"agro-mart/internal/repository"
)

type UserService interface {
	GetAllUser(id string) (*model.User, error)
	UpsertUser(*model.User) error
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (impl UserServiceImpl) GetAllUser(id string) (*model.User, error) {
	user, err := impl.UserRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	userModel := mapper.UserDomainToUserModel(user)

	return userModel, nil
}

func (impl UserServiceImpl) UpsertUser(user *model.User) error {
	userDomain := mapper.UserModelToUserDomain(user)

	err := impl.UserRepository.UpsertUser(userDomain)
	if err != nil {
		return err
	}

	return nil
}
