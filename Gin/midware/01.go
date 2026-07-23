// 中间件
package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default() // 默认使用了 Logger 和 Recovery 中间件
	// // recovery 中间件用于捕获 panic 并返回 500 状态码 出现panic 就中断当前 并返回500
	// r := gin.New() // 不使用任何中间件  创建一个空白引擎

}

func Test() {
	r := gin.New() // 空白engin

	//全局使用
	r.Use(gin.Logger())

	// 分组
	v1 := r.Group("/v1")
	{
		v1.Use(gin.Logger())
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "test"})
		})
	}
	// 单独使用
	r.GET("/test", gin.Logger(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// 错误处理中间件
//
//	在主 函数 hanlder执行完之后 执行的错误处理函数
func errorHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next() //

		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors.ByType(gin.ErrorTypeAny).String()})
		}
	}

}

func test02() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(errorHandler())
	r.GET("/ok", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	r.GET("/error", func(c *gin.Context) {
		c.Error(errors.New("test error"))
	})
	r.Run(":8080")
}

// basicAuth

// HTTP 基本认证以 Base64 编码字符串传输凭据，未加密。任何能拦截流量的人都可以轻易解码凭据
// 。在生产环境中使用 BasicAuth 时务必使用 HTTPS（TLS）。

// simulate some private data
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	router := gin.Default()

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
