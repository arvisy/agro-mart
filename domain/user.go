package domain

type User struct {
	Id   *string
	Name *string
	Role *string
	Audit
}

func (User) TableName() string {
	return "user"
}
