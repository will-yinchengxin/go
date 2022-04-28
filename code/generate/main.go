package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{ // H是一个开箱即用的map
			"message": "Hello world!",
		})
	})
	r.Run(":8080")
}
