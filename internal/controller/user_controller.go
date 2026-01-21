package controller

import (
	"agro-mart/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func (impl UserController) GetAllUser(c *gin.Context) {
	res, err := impl.UserService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
