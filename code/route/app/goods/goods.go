package goods

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
	fmt.Println("getGoods ===>>> start")
	c.JSON(200, gin.H{
		"name":"goods",
		"type":"normal",
	})
	fmt.Println("getGoods ===>>> end")
}