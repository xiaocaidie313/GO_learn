# IM-system

## 要点

- TCP 聊天室：Server 监听，Client 主动连接
- **User** = 服务端视角的一个在线连接（不是 Client）
- **Client** = 用户电脑上的聊天程序（替代 nc）
- 消息经 channel 在服务端协程间分发

## 架构

```
Client.conn ←→ User.conn
Server: OlineMap + Message chan + MessageListener
User:   C chan + Listener + DoMessage
```

## 消息协议

| 命令 | 含义 |
|------|------|
| 普通文本 | 公聊广播 |
| `who` | 查在线用户 |
| `rename\|新名` | 改名 |
| `to\|用户名\|内容` | 私聊 |

## 关键 API

| API | 说明 |
|-----|------|
| `net.Listen` / `net.Accept` | 服务端 |
| `net.Dial` | 客户端连接 |
| `conn.Read` / `conn.Write` | 收发字节 |
| `BoardCast` | 写入 Message chan |
| `select` + `time.After` | 30 秒无消息踢人 |

## 示例文件

- `main.go` — 启动 Server
- `server.go` — Handler、广播、超时
- `user.go` — 上线/下线、DoMessage
- `client.go` — 菜单、公聊/私聊

## 运行

```powershell
# 终端1：服务端
cd IM-system
go run .

# 终端2：客户端
go run client.go -ip 127.0.0.1 -port 8888
```
