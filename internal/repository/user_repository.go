package repository

import (
	"agro-mart/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]domain.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (impl UserRepositoryImpl) GetAllUser() ([]domain.User, error) {
	q := impl.DB.Debug().Model(domain.User{})

	users := []domain.User{}
	tx := q.Find(users)
	if tx.Error != nil {
		return []domain.User{}, tx.Error
	}

	return users, nil
}
