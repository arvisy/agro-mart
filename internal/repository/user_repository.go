package repository

import (
	"agro-mart/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(id string) (*domain.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (impl UserRepositoryImpl) GetUserById(id string) (*domain.User, error) {
	q := impl.DB.Debug().Model(domain.User{})

	user := domain.User{}
	tx := q.Where("id = ?", id).Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
