package gin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
	"mime/multipart"
)

// 上传文件
type UploadPartCreate struct {
	FileName   string                `json:"fileName" form:"fileName" validate:"required"`
	Path       string                `json:"path" form:"path" validate:"required"`
	UploadID   string                `json:"uploadId" form:"uploadId" validate:"required"`
	PartNumber int                   `json:"partNumber" form:"partNumber" validate:"required"`
	UploadData *multipart.FileHeader `json:"uploadData" form:"uploadData" validate:"required"`
}

type User struct {
	Name    string
	Habits []string
}



// 中间件 middleware 在 Golang 中是一个很重要的概念，与 Java 中的拦截器类似
// 我们是直接通过 gin.Default() 来初始化 gin 对象,其中它包含了一个自带默认中间件的 *Engine
// 其中，Default() 函数会默认绑定两个已经准备好的中间件，它们就是 Logger 和 Recovery，帮助我们打印日志输出和 painc 处理。
// 从上面 Default()函数中，我们可以看到 Gin 中间件是通过 Use 方法设置的，它是一个可变参数，可同时设置多个中间件
// 设置中间件需要满足两个条件: 1)它是个函数 2) 函数的返回类型必须是一个 HandlerFunc
func costTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		nowTime := time.Now()
		costTime := time.Since(nowTime)
		requestUrl := c.Request.URL.String()
		fmt.Printf("the request url %s, cost time is %v \n", requestUrl, costTime)
		// 处理请求
		c.Next()
	}
}

func Start() {
	gin.SetMode("release")
	route := gin.Default()
	// 注册中间件
	route.Use(costTimeMiddleware())
	route.GET("/test", func(c *gin.Context) {
		// 查看请求参数
		fmt.Println(c.Request.URL.Path)
		fmt.Println(c.Request.URL)
		fmt.Println(c.Request.Header)
		fmt.Println(c.Request.Host)
		fmt.Println(c.Request.Proto)

		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	route.Run(":8080")
}

func write(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Custom-Header", "custom")
	w.WriteHeader(201)
	user := &User {
		Name:    "aoho",
		Habits: []string{"balls", "running", "hiking"},
	}
	json, _ := json.Marshal(user)
	w.Write(json)
}

func init() {
	//rand.Seed(time.Now().Unix())
}

/*
编写 HTTPS 服务器
HTTPS = HTTP + Secure(安全)

RSA 进行加密
SHA 进行验证
密钥和证书

生成密钥文件
openssl genrsa -out /d/projectserver.key 2048

生成证书文件
openssl req -new -x509 -sha256 -key server.key -out /d/project/server.crt -days 3650

*/
func Https() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("path:", r.URL.Path)
		fmt.Println("Url:",r.URL)
		fmt.Println("Host:",r.Host)
		fmt.Println("Header:",r.Header)
		fmt.Println("Method:",r.Method)
		fmt.Println("Proto:",r.Proto)
		fmt.Println("UserAgent:",r.UserAgent())

		scheme := "http://"
		if r.TLS != nil {
			scheme = "https://"
		}
		fmt.Println("完整的请求路径:", strings.Join([]string{scheme,r.Host,r.RequestURI},""))
		fmt.Fprintf(w, "Hello Go Web")
	})

	fmt.Println("HTTPS 服务器已经启动，请在浏览器地址栏中输入 https://localhost:8080/")

	err := http.ListenAndServeTLS(":8080","D:\\Project\\server.crt","D:\\Project\\server.key",nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}
}
