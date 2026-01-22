package controller

import (
	"agro-mart/internal/model"
	"agro-mart/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	UserService service.UserService
}

func (impl UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id format",
		})
		return
	}

	res, err := impl.UserService.GetAllUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (impl UserController) UpsertUser(c *gin.Context) {
	var bodyReq model.User
	err := c.ShouldBindJSON(&bodyReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = impl.UserService.UpsertUser(&bodyReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
