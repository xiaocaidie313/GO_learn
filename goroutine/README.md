# goroutine

## 要点

- `go f()` 启动新协程，与主协程并发执行
- **主 goroutine 结束，整个程序退出**，其他协程也会被终止
- 匿名函数：`go func() { ... }()`
- `runtime.Goexit()` 终止**当前** goroutine（不影响调用它的 goroutine）

## 常用方法

| 操作 | 语法 |
|------|------|
| 启动协程 | `go funcName()` |
| 匿名协程 | `go func() { ... }()` |
| 退出当前协程 | `runtime.Goexit()` |
| 休眠 | `time.Sleep(d)` |

## 示例文件

- `01.go` — 主协程与子协程并发打印
- `02.go` — 嵌套 goroutine 与 Goexit

## 运行

```powershell
go run ./goroutine/01.go
go run ./goroutine/02.go
```
