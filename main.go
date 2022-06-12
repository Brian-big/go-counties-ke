package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type County struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	HeadQuarters string `json:"headquarters"`
}

var counties = []County{
	{ID: "1", Name: "Mombasa", HeadQuarters: "Mombasa city"},
	{ID: "2", Name: "Kwale", HeadQuarters: "Kwale"},
	{ID: "3", Name: "Kilifi", HeadQuarters: "Kilifi"},
	{ID: "4", Name: "Tana River", HeadQuarters: "Hola"},
}

func main() {
	fmt.Println("Go Kenyan Counties - Know Kenya movement")
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go Kenyan Counties - Know Kenya movement",
		})
	})

	r.GET("/counties", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, counties)
	})

	r.Run()

}
