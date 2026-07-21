package main

/*
突然有个新感悟 后端部分 是处理请求 + 拿数据 + 返回请求的。所以有回调函数。 回调函数部分就是 处理 返回的请求的样子是怎么样

*/

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	setupUpload(r)
	setupRedirect(r)
	setupArticles(r)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run(":8080")
}

// 1 路径参数 // :name  /*action

/*
    1 :name —— 匹配单个路径段。例如，/user/:name 匹配 /user/john, 但不匹配 /user/ 或 /user。
    2 *action —— 匹配前缀之后的所有内容，包括斜杠。
例如，/user/:name/*action 匹配 /user/john/send 和 /user/john/。捕获的值包含前导 /。
*/

// 查询字符串参数

// 1:  c.query()  2  .c.defaultquery()
/*
	查询字符串参数是出现在 URL 中 ? 后面的键值对（例如 /search?q=gin&page=2）。Gin 提供了两种方法来读取它们：

	c.Query("key") 返回查询参数的值，如果键不存在则返回空字符串。
	c.DefaultQuery("key", "default") 返回值，如果键不存在则返回指定的默认值。

*/

// 表单

// 1:  c.postform()  2  .c.defaultpostform()
/*
	表单参数是出现在请求体中的键值对（例如 name=John&age=20）。Gin 提供了两种方法来读取它们：

	c.PostForm("key") 返回表单参数的值，如果键不存在则返回空字符串。
	c.DefaultPostForm("key", "default") 返回值，如果键不存在则返回指定的默认值。

	适用于 application/x-www-form-urlencoded 和 multipart/form-data 内容类型

*/


// 解析 map
/*
	c.QueryMap 和 c.PostFormMap 来将方括号表示法的参数（如 ids[a]=1234）解析为 map[string]string。

	c.QueryMap("key") —— 从 URL 查询字符串中解析 key[subkey]=value 形式的键值对。
	c.PostFormMap("key") —— 从请求体中解析 key[subkey]=value 形式的键值对
*/
// 文件上床

/*
	c.FormFile("fieldname") —— 从 multipart/form-data 请求中获取文件。
	c.SaveUploadedFile(file *multipart.FileHeader, dst string) —— 将上传的文件保存到指定路径。
*/

func setupUpload(router *gin.Engine) {
	router.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["file"]

		_ = os.MkdirAll("./uploads/", 0755)
		for _, file := range files {
			dst := filepath.Join("./uploads/", file.Filename)
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("save err: %s", err.Error()))
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "count": len(files)})
	})
}

func setupRedirect(r *gin.Engine) {
	r.GET("/oldone", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
}

// page
func setupArticles(r *gin.Engine) {
	r.GET("/api/articles", func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
		offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

		if limit > 100 {
			limit = 100 // cap the page size
		}

		// articles, total := db.ListArticles(limit, offset)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []gin.H{}, // articles
			"meta": gin.H{
				"limit":  limit,
				"offset": offset,
				"total":  0, // total
			},
		})
	})
}

// 原 TESTMAIN / TESTMAIN2 逻辑已分别抽到 setupUpload / setupRedirect，由 main 统一注册
