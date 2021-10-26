package main

import (
	"github.com/gin-gonic/gin"
	"code/status"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(context *gin.Context) {
		status.ResponseError(context, status.CodeSucces)
	})
	r.Run(":8001")
}
