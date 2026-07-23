// All rendering methods share this pattern:
// c.JSON(http.StatusOK, data)   // application/json
// c.XML(http.StatusOK, data)    // application/xml
// c.YAML(http.StatusOK, data)   // application/x-yaml
// c.TOML(http.StatusOK, data)   // application/toml
// c.ProtoBuf(http.StatusOK, data) // application/x-protobuf

// 所有渲染方法都接受一个 HTTP 状态码和一个数据值。Gin 会序列化数据并自动设置适当的 Content-Type 头。

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"name"`
			Content string `json:"content"`
		}
		msg.Name = "Alex"
		msg.Content = "Hello, World!"
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
		})

	})

}

// secureJSON

// You can also use your own secure json prefix
// router.SecureJsonPrefix(")]}',\n")
// Will output  :   while(1);["lena","austin","foo"]
// c.SecureJSON(http.StatusOK, names)
/*
当响应数据是 JSON 数组时，SecureJSON 在响应体前添加一个不可解析的前缀——默认为 while(1);。
这会导致浏览器的 JavaScript 引擎进入无限循环（如果通过 <script> 标签加载响应），
从而防止数据被访问。合法的 API 消费者（使用 fetch、XMLHttpRequest 或任何 HTTP 客户端）读取原始响应体，
只需在解析前去除前缀即可。

Google 的 API 使用类似技术 )]}'\n，Facebook 使用 for(;;);。
你可以通过 router.SecureJsonPrefix() 自定义前缀

*/

// PureJson
// 当你的 API 消费者期望原始的、未转义的 JSON 时使用 PureJSON。当响应可能嵌入 HTML 页面时使用标准的 JSON。
// c.PureJSON(http.StatusOK, data)

// static

/*
router.Static(relativePath, root) — 提供整个目录。对 relativePath 的请求会映射到 root 下的文件。例如，router.Static("/assets", "./assets") 会在 /assets/style.css 处提供 ./assets/style.css。
router.StaticFS(relativePath, fs) — 类似于 Static，但接受一个 http.FileSystem 接口，让你可以更好地控制文件解析方式。当你需要从嵌入式文件系统提供文件或自定义目录列表行为时使用。
router.StaticFile(relativePath, filePath) — 提供单个文件。适用于 /favicon.ico 或 /robots.txt 等端点。


*/

// c.File() 返回本地的某个文件
/* router.GET("/local/file", func(c *gin.Context) {
  c.File("local/file.go")
})
*/
func FileTest() {
	router := gin.Default()

	// Serve a file inline (displayed in browser)
	router.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})

	// Serve a file from an http.FileSystem
	var fs http.FileSystem = http.Dir("/var/www/assets")
	router.GET("/fs/file", func(c *gin.Context) {
		c.FileFromFS("fs/file.go", fs)
	})

	// Serve a file as a downloadable attachment with a custom filename
	router.GET("/download", func(c *gin.Context) {
		c.FileAttachment("local/report-2024-q1.xlsx", "quarterly-report.xlsx")
	})
}

// c.File(path) — 从本地文件系统提供文件。内容类型会自动检测。
// 当你在编译时已知确切的文件路径或已验证过路径时使用。
// c.FileFromFS(path, fs) — 从 http.FileSystem 接口提供文件。
// 适用于从嵌入式文件系统（embed.FS）、自定义存储后端提供文件，或当你想限制对特定目录树的访问时
// 。
// c.FileAttachment(path, filename) — 通过设置 Content-Disposition: attachment 头将文件作为下载提供。
// 浏览器会提示用户使用你提供的文件名保存文件，而不管磁盘上的原始文件名。
