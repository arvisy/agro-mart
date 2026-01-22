package config

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine, c *Container) {
	api := r.Group("api/v1")
	{
		api.GET("/user/:id", c.userController.GetUserById)
		api.POST("/user", c.userController.UpsertUser)
		api.GET("/product", c.productController.GetAllProduct)
	}
}
