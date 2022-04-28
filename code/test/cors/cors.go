package cors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// X-Forwarded-For: 追加的是请求方的 IP 地址
		// X-Real-IP
		keys := "x-token,token, X-Token,Authorization,content-type"
		for k, _ := range c.Request.Header {
			if keys != "" {
				keys = keys + ", " + k
			} else {
				keys += k
			}
		}
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", keys)
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
