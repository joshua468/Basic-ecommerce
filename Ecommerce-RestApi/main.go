package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/rssagg/Ecommerce-RestApi/routes"
)

func main() {
	router := gin.Default()
	routes.InitializeRoutes(router)

	router.Run(":8080")
}
