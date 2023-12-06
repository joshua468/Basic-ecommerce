package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/rssagg/Ecommerce-RestApi/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/products", handlers.GetProduct)
	router.GET("/products/:id", handlers.GetProductById)
	router.GET("/products", handlers.CreateProducts)
}
