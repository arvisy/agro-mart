package user

import "time"

type User struct {
	Id        *string
	Name      *string
	Role      *string
	CreatedAt *time.Time
}

func (User) TableName() string {
	return "user"
}
