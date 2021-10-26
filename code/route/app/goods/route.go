package goods

import (
	"gin/frame/route"
	"github.com/gin-gonic/gin"
)

func init() { // 初始化的时候注册
	route.RegisterRoute(Routes)
}

func Routes(g *gin.Engine) {
	g.GET("/getGoods", GetGoods)
}