# databind

Gin 数据绑定：把请求里的参数自动填进结构体，并可做校验。

## 要点

```
请求参数（Query/Form/JSON/URI/Header）
        ↓ ShouldBind / ShouldBindXxx
结构体字段（按类型转换 + binding 校验）
        ↓
业务代码使用结构体
```

- **类型**看结构体字段声明；`ShouldBind` 负责转换，不会自己猜类型
- **Must Bind**（`Bind`）：失败直接 400 中止，你不好再改状态码
- **Should Bind**（`ShouldBind`）：失败返回 `error`，自己处理（推荐）
- `ShouldBind` 内部最终会走到 `ShouldBindWith(obj, binding.JSON/Form/...)`

## 常用方法

| 方法 | 说明 |
|------|------|
| `c.ShouldBind(&obj)` | 按 Content-Type 自动选绑定器 |
| `c.ShouldBindJSON(&obj)` | 固定绑 JSON |
| `c.ShouldBindQuery(&obj)` | 固定绑 Query |
| `c.ShouldBindUri(&obj)` | 绑路径参数 `:name` |
| `c.ShouldBindHeader(&obj)` | 绑请求头 |
| `c.ShouldBindWith(&obj, binding.Query)` | 自己指定绑定器 |
| `c.BindXxx` | Must 系列，失败直接 Abort 400 |

## struct tag

| tag | 用途 |
|-----|------|
| `form:"name"` | Query / Form 字段名 |
| `form:"name,default=Tom"` | 默认值 |
| `json:"name"` | JSON 字段名 |
| `uri:"name"` | 路径参数 |
| `header:"Rate"` | 请求头 |
| `binding:"required"` | 校验规则 |
| `time_format:"2006-01-02"` | 时间解析格式 |
| `collection_format:"csv"` | 集合拆分方式 |
| `parser=encoding.TextUnmarshaler` | 自定义解析 |

## 文件

| 文件 | 内容 | 运行 |
|------|------|------|
| `01.go` | 自定义校验 bookabledate + Query 绑定日期 | `go run 01.go` |
| `02.go` | ShouldBind / URI / 默认值 tag | `go run 02.go` |
| `03.go` | TextUnmarshaler 自定义解析 + Header 绑定 | `go run 03.go` |

## 测试示例

```powershell
cd C:\Users\Liam\Desktop\go\Gin\databind

go run 01.go
# http://localhost:8085/bookable?check_in=2026-08-01&check_out=2026-08-05

go run 02.go
# http://localhost:8080/testing?name=Tom&age=20
# http://localhost:8080/uri/Tom/20

go run 03.go
# http://localhost:8088/test?birthday=2020-09-01&birthdays=2020-09-01,2020-09-02
```
