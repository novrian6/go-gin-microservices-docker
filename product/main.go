package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/products", getProductsHandler)
	log.Println("Product service started on port 8081")
	r.Run(":8081")
}

func getProductsHandler(c *gin.Context) {
	// Simulate fetching products from a database
	c.JSON(http.StatusOK, gin.H{"message": "List of products"})
}
