# Gin

Go 最流行的 Web 框架之一，用于处理 HTTP 请求并返回响应。

## 要点

```
后端 = 处理请求 + 拿数据 + 返回响应
      回调函数（handler）就是「怎么处理、返回什么」
```

## 常用方法

| 类别 | 方法 | 说明 |
|------|------|------|
| 路由 | `r.GET/POST/PUT/DELETE(path, handler)` | 注册路由 |
| 路径参数 | `c.Param("name")` | `/user/:name` |
| 通配 | `c.Param("action")` | `/user/:name/*action` |
| 查询参数 | `c.Query("key")` | URL `?key=val` |
| 查询默认值 | `c.DefaultQuery("key", "default")` | 不存在时用默认值 |
| 表单 | `c.PostForm("key")` | POST body 表单 |
| 表单默认 | `c.DefaultPostForm("key", "default")` | 表单默认值 |
| Map 解析 | `c.QueryMap("key")` / `c.PostFormMap("key")` | `ids[a]=1234` |
| 文件上传 | `c.FormFile("field")` | multipart 单文件 |
| 保存文件 | `c.SaveUploadedFile(file, dst)` | 存到磁盘 |
| 响应 JSON | `c.JSON(code, obj)` | 返回 JSON |
| 重定向 | `c.Redirect(code, url)` | 301/302 跳转 |

## 路径参数说明

- `:name` — 匹配**单个**路径段，如 `/user/john`
- `*action` — 匹配前缀后**所有内容**（含 `/`），如 `/user/john/send`

## 示例文件

- `router.go` — ping、上传、重定向、分页查询 API

## 运行

```powershell
cd C:\Users\Liam\Desktop\go
go run ./Gin/

# 测试
curl http://localhost:8080/ping
curl "http://localhost:8080/api/articles?limit=10&offset=0"
```

上传目录需存在 `./uploads/`（程序会自动创建）。
