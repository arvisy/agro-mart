package user

type User struct {
	Id   *string
	Name *string
	Role *string
}

func (User) TableName() string {
	return "user"
}
