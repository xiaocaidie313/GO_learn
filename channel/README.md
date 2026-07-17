# channel

## 要点

- channel 是 goroutine 之间通信的管道
- **无缓冲** channel：发送和接收必须同时就绪，否则阻塞
- **有缓冲** channel：`make(chan T, n)`，缓冲区满/空时才阻塞
- 关闭 channel 用 `close(c)`，接收方可用 `for v := range c` 读至结束
- `select` 可同时监听多个 channel

## 常用方法

| 操作 | 语法 | 说明 |
|------|------|------|
| 创建 | `make(chan int)` | 无缓冲 |
| 创建 | `make(chan int, 3)` | 缓冲容量 3 |
| 发送 | `ch <- v` | 写入 |
| 接收 | `v := <-ch` | 读出 |
| 关闭 | `close(ch)` | 关闭后不能再写 |
| 多路监听 | `select { case ... }` | 类似 switch |

## 示例文件

- `01.go` — 无缓冲、有缓冲、select 示例

## 运行

```powershell
go run ./channel/
```
