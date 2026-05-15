package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sre/internal/service"
)

func GetProductsHandler(c *gin.Context) {

	products, err := service.GetProducts()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get products",
		})

		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductByIDHandler(c *gin.Context) {

	id := c.Param("id")

	product, err := service.GetProduct(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})

		return
	}

	c.JSON(http.StatusOK, product)
}
