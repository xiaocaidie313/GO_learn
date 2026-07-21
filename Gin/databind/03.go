package main

import (
	"encoding"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Birthday string

/*
 要覆盖 Gin 的默认绑定逻辑，请在你的类型上定义一个满足
 Go 标准库中 encoding.TextUnmarshaler 接口的函数。
 然后在要绑定的字段的 uri/form 标签中指定 parser=encoding.TextUnmarshaler。
*/

func (b *Birthday) UnmarshalText(text []byte) error {
	*b = Birthday(strings.Replace(string(text), "-", "/", -1))
	return nil
}

// encoding 的 实现 需要在tag里写上 parser=encoding.TextUnmarshaler
// 如果是 binding.BindUnmarshaler 则不需要写上
var _ encoding.TextUnmarshaler = (*Birthday)(nil)

type testHeader struct {
	Rate  int `header:"Rate"`
	Limit int `header:"Limit"`
}

type Rate struct {
	Rate int `header:"Rate"`
}

type Limit struct {
	Limit int `header:"Limit"`
}

func main() {
	route := gin.Default()

	// TextUnmarshaler 自定义解析：2020-09-01 → 2020/09/01
	route.GET("/test", func(ctx *gin.Context) {
		var request struct {
			Birthday         Birthday   `form:"birthday,parser=encoding.TextUnmarshaler"`
			Birthdays        []Birthday `form:"birthdays,parser=encoding.TextUnmarshaler" collection_format:"csv"`
			BirthdaysDefault []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02,parser=encoding.TextUnmarshaler" collection_format:"csv"`
		}
		if err := ctx.ShouldBindQuery(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, request)
	})

	// 绑定 Header 请求头
	route.GET("/header", func(c *gin.Context) {
		var h testHeader
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, h)
	})

	_ = route.Run(":8088")
}

// 绑定 Header 请求头
/*
func (c *Context){
	var h header
	if err := c.ShouldBindHeader(&h); err != nil {
		return err
	}
	return nil
}
*/
