package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 当你需要轻松撤销时使用会话（例如注销、封禁用户）。当你需要跨微服务的无状态认证时使用 JWT。

// 可以有不同的后端 redis  cookie 等作为这里的store
func main() {
	r := gin.Default()

	// Create cookie-based session store with a secret key
	store := cookie.NewStore([]byte("your-secret-key")) // 这里的后端用 cookie 存储
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user", "john") //  set 内部实现是 map[key] value 的键值对
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "logged in"})
	})

	r.GET("/profile", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "logged out"})
	})

	r.Run(":8080")
}
