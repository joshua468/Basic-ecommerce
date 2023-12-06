package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/rssagg/Ecommerce-RestApi/models"
)

var products []models.Product

func GetProduct(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
}

func CreateProducts(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	c.JSON(http.StatusCreated, newProduct)
}
