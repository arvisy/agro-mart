package repository

import (
	"agro-mart/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	GetUserById(id string) (*domain.User, error)
	UpsertUser(user *domain.User) error
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

func (impl UserRepositoryImpl) UpsertUser(user *domain.User) error {
	q := impl.DB.Debug().Model(domain.User{})

	tx := q.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "role"}),
	}).Create(user)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
