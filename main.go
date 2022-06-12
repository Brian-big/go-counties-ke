package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Go Kenyan Counties - Know Kenya movement")
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go Kenyan Counties - Know Kenya movement",
		})
	})
	r.Run()

}
