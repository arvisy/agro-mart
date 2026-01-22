package controller

import (
	"agro-mart/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService service.ProductService
}

func (impl ProductController) GetAllProduct(c *gin.Context) {
	res, err := impl.ProductService.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
