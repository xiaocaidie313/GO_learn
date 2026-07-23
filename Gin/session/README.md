# session

Gin Session：在多次请求之间保存登录态等数据。

## 要点

```
需要轻松撤销（注销、封禁）→ Session
跨微服务无状态认证     → JWT
```

- Store 可换后端：`cookie` / `redis` / `memory` 等
- Cookie Store：数据加密后放浏览器 Cookie，密钥用于防篡改
- `Sessions` 中间件挂上后，handler 里用 `sessions.Default(c)`

## 常用方法

| 方法 | 说明 |
|------|------|
| `cookie.NewStore(secret)` | Cookie 后端，secret 用于签名/加密 |
| `sessions.Sessions(name, store)` | 注册 Session 中间件，`name` 是 Cookie 名 |
| `sessions.Default(c)` | 取当前请求的 Session |
| `session.Set(k, v)` | 写入（内部类似 map） |
| `session.Get(k)` | 读取，没有则 `nil` |
| `session.Clear()` | 清空 |
| `session.Save()` | 写回响应（务必调用） |

## 文件

- `01.go` — Cookie Session：登录 / 个人页 / 注销

## 运行

```powershell
cd C:\Users\Liam\Desktop\go
go get github.com/gin-contrib/sessions
go get github.com/gin-contrib/sessions/cookie
go run ./Gin/session/01.go
```

```text
GET /login    → 写入 user
GET /profile  → 读取 user
GET /logout   → 清空
```
