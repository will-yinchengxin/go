package route

import (
	"github.com/gin-gonic/gin"
)

// 自定义注册路由类型
type Router func(engine *gin.Engine)

// 创建切片存储路由
var routers = []Router{}

func RegisterRoute(routes ...Router) {
	routers = append(routers, routes...)
}

func InitRouter() *gin.Engine {
	//r := gin.Default()
	//// GET：请求方式；/hello：请求的路径
	//// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	//r.GET("/hello", func(c *gin.Context) {
	//	// c.JSON：返回JSON格式的数据
	//	c.JSON(200, gin.H{
	//		"message": "Hello world!",
	//	})
	//})
	//// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	//r.Run(":8060")

	r := gin.Default()

	for _, route := range routers {
		route(r)
	}

	return r
}
