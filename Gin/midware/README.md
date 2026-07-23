# midware（中间件）

中间件包住业务 handler：请求前 / 请求后插入逻辑，类似 Express 的 `use + next`。

## 要点

```
请求 → Mid1 前 → Mid2 前 → Handler → Mid2 后 → Mid1 后 → 响应
```

- `c.Next()`：进入下一层；没有 Next → handler 不执行
- `c.Abort()`：打断后续链
- `gin.Default()` = `New()` + Logger + Recovery
- 协程里不要直接用 `c`，要用 `c.Copy()` 或拷出的值

## 常用方法

| 方法 | 说明 |
|------|------|
| `r.Use(mid)` | 全局中间件 |
| `g.Use(mid)` | 路由组中间件 |
| `r.GET(path, mid, handler)` | 单路由中间件 |
| `c.Next()` | 调用后续 handler |
| `c.Abort()` / `AbortWithStatusJSON` | 中断 |
| `c.Set` / `c.Get` / `MustGet` | 在链上传请求级数据 |
| `c.Error(err)` | 记录错误，供错误中间件处理 |
| `gin.BasicAuth(accounts)` | HTTP Basic 认证 |
| `gin.AuthUserKey` | BasicAuth 写入的用户名 key |

## 文件

| 文件 | 内容 |
|------|------|
| `01.go` | Logger、错误处理、BasicAuth 可运行示例 |
| `02.md` | 安全响应头说明 + 示例代码 |

## 运行

```powershell
go run ./Gin/midware/01.go
# GET /test
# GET /ok  /error
# GET /admin/secrets  （浏览器弹账号密码，如 foo/bar）
```
