package config

import (
	"agro-mart/internal/controller"
	"agro-mart/internal/repository"
	"agro-mart/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	DB *gorm.DB

	userController controller.UserController
}

func NewContainer(db *gorm.DB) *Container {
	c := &Container{DB: db}
	c.initController()

	return c
}

func (c *Container) initController() {
	c.userController = controller.UserController{
		UserService: service.UserServiceImpl{
			UserRepository: repository.UserRepositoryImpl{DB: c.DB},
		},
	}
}
