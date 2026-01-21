package config

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine, c *Container) {
	api := r.Group("api/v1")
	{
		api.GET("/users", c.userController.GetAllUser)
	}
}
