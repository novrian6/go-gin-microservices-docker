package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", getUsersHandler)
	log.Println("User service started on port 8080")
	r.Run(":8080")
}

func getUsersHandler(c *gin.Context) {
	// Simulate fetching users from a database
	c.JSON(http.StatusOK, gin.H{"message": "List of users"})
}
