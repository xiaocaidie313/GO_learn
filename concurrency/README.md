# concurrency

Go 并发编程核心：goroutine + channel + sync + select + context。

## 文件一览

| 文件 | 主题 | 运行 |
|------|------|------|
| `coroutine.go` | goroutine 启动、主协程退出问题 | `go run coroutine.go` |
| `channel.go` | 无缓冲/有缓冲 channel | `go run channel.go` |
| `waitgroup.go` | WaitGroup 等待一组协程 | `go run waitgroup.go` |
| `select.go` | select 多路监听、永久阻塞 | `go run select.go` |
| `context.go` | Context 取消/超时/传值 | `go run context.go` |
| `sync.go` | Mutex / Once / Pool | `go run sync.go` |

## 核心概念对比

| 工具 | 干什么 | 类比 |
|------|--------|------|
| **goroutine** | 轻量并发 | 开一个后台任务 |
| **channel** | 协程间传数据、同步 | 管道/队列 |
| **WaitGroup** | 等 N 个协程全部完成 | 计数器：`Add/Done/Wait` |
| **select** | 同时等多个 channel/超时 | 单片机主循环等多事件 |
| **Context** | 取消、超时、跨层传元数据 | 任务控制面板 |
| **Mutex** | 保护共享变量 | 互斥锁 |
| **sync.Once** | 只执行一次 | 单例初始化 |
| **sync.Pool** | 临时对象复用 | 对象池（可被 GC 清空） |

## select{} vs for{}

```go
for {}      // 死循环，占 CPU
select {}   // 永久阻塞，几乎不占 CPU（类似单片机 WFI 等中断）
```

## Context 类型（标准库内部分层）

```
Background (emptyCtx)
  → WithValue  → valueCtx     带 key-value
  → WithCancel → cancelCtx    可 cancel
  → WithTimeout → timerCtx    带 deadline
```

Context 不是 handler，而是「一次任务的生存环境和控制信号」。

## 常用 API 速查

```go
go f()                              // 启动协程
ch := make(chan int)                // 无缓冲
ch := make(chan int, 3)             // 缓冲 3
close(ch)                           // 关闭 channel

var wg sync.WaitGroup
wg.Add(1); go func(){ defer wg.Done(); ... }(); wg.Wait()

select { case v := <-ch: case ch <- v: case <-time.After(d): }

ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
select { case <-ctx.Done(): return }

var mu sync.Mutex
mu.Lock(); defer mu.Unlock()
```

## 注意

- 各 `.go` 文件中保留了学习注释，阻塞/报错示例以注释形式保留
- 主 goroutine 退出 → 整个程序退出，其他协程也会被终止
- 无缓冲 channel 发送/接收必须同时就绪，否则会阻塞
- nil channel 在 select 中会被忽略；单独读写 nil channel 会永久阻塞
- `sync.Pool` 只适合**临时对象**，不适合数据库连接等长生命周期资源
