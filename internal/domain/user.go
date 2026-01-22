package domain

import (
	"agro-mart/internal/enum"
	"time"
)

type User struct {
	Id        *string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      *string
	Role      *enum.UserRole
	CreatedAt *time.Time
}

func (User) TableName() string {
	return "user"
}
