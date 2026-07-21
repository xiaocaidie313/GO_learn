/*
	类型 - Must bind
方法 - Bind、BindJSON、BindXML、BindQuery、BindYAML
行为 - 这些方法底层使用 MustBindWith。如果存在绑定错误，请求将使用 c.AbortWithError(400, err).SetType(ErrorTypeBind) 中止。这会将响应状态码设置为 400，并将 Content-Type 头设置为 text/plain; charset=utf-8。注意，如果你在此之后尝试设置响应码，将会出现警告 [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422。如果你希望更好地控制行为，请考虑使用 ShouldBind 等效方法。
	类型 - Should bind
方法 - ShouldBind、ShouldBindJSON、ShouldBindXML、ShouldBindQuery、ShouldBindYAML
行为 - 这些方法底层使用 ShouldBindWith。如果存在绑定错误，错误会被返回，由开发者负责适当地处理请求和错误。

自己管理err

*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name" binding:"required"`
	Age  int    `form:"age" binding:"required"`
}

type Dog struct {
	Name string `uri:"name" binding:"required"`
	Age  int    `uri:"age" binding:"required"`
}

func main() {
	r := gin.Default()

	// Query / Form：ShouldBind（自己管理 err）
	r.GET("/testing", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "OK", "name": person.Name, "age": person.Age})
	})

	// 路径参数：用 uri tag + ShouldBindUri
	r.GET("/uri/:name/:age", func(c *gin.Context) {
		var dog Dog
		if err := c.ShouldBindUri(&dog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "OK", "name": dog.Name, "age": dog.Age})
	})

	// 默认值示例
	r.GET("/default", func(c *gin.Context) {
		var p Person2
		_ = c.ShouldBind(&p)
		c.JSON(http.StatusOK, p)
	})

	r.Run(":8080")
}

/*
将默认值放在表单键之后：form:"name,default=William"。
对于集合，使用 collection_format:"multi|csv|ssv|tsv|pipes" 指定如何拆分值。
对于 multi 和 csv，在默认值中使用分号分隔值（例如 default=1;2;3）。Gin 会在内部将其转换为逗号，以保持标签解析器的明确性。
对于 ssv（空格）、tsv（制表符）和 pipes（|），在默认值中使用自然分隔符。
*/

// /
type Person2 struct {
	Name      string    `form:"name,default=William"`
	Age       int       `form:"age,default=10"`
	Friends   []string  `form:"friends,default=Will;Bill"` // multi/csv: use ; in defaults
	Addresses [2]string `form:"addresses,default=foo bar" collection_format:"ssv"`
	LapTimes  []int     `form:"lap_times,default=1;2;3" collection_format:"csv"`
}

// BindUnmarshaler 自定义 反序列化器
