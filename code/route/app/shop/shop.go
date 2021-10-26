package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetShops(c *gin.Context) {
	fmt.Println("getShop ===>>> start")
	c.JSON(200, gin.H{
		"name":"shop",
		"type":"normal",
	})
	fmt.Println("getShop ===>>> end")
}
