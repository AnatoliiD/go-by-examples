package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Item is an awesome json model 🤘
type Item struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

var items []Item

func init() {
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("Awesome Item#%d", i)
		fmt.Println(name)
		items = append(items, Item{
			Name:   name,
			Number: i,
		})
	}
}

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("static", false)))
	r.StaticFile("/", "static/index.html")
	api := r.Group("/api")
	api.GET("/items", func(c *gin.Context) {
		c.JSON(200, &items)
	})
	r.Run(":" + os.Getenv("PORT"))
}
