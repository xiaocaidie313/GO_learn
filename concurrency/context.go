package main

import (
	"context"
	"fmt"
	"time"
)

// Context是一个接口，定义了四个方法
// 相比于管道和 WaitGroup，它可以更好的控制子孙协程以及层级更深的协程
//
// type Context interface {
// 	Deadline() (deadline time.Time, ok bool)
// 	Done() <-chan struct{}
// 	Err() error
// 	Value(key any) any
// }
//
// emptyCtx 是 context 包中的一个结构体，用于 提供根context 和 占位context
// Background() / TODO() 返回 emptyCtx
// emptyCtx 的这些方法都实现不了 所以返回的都是空值
//
// valueCtx 是 context 包中的一个结构体，用于 提供value context
// type valueCtx struct {
// 	Context       // 父context
// 	key, val any
// }
// func (c *valueCtx) Value(key any) any {
// 	if c.key == key {
// 		return c.val
// 	}
// 	return value(c.Context, key) // 递归调用，向上找父 context
// }
//
// 标准库分层：Background → WithValue → WithCancel → WithTimeout

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("  %s 收到取消: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("  %s 工作中...\n", name)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("=== WithCancel ===")
	ctx1, cancel := context.WithCancel(context.Background())
	go worker(ctx1, "worker-A")
	time.Sleep(700 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== WithTimeout ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel2()
	go worker(ctx2, "worker-B")
	time.Sleep(800 * time.Millisecond)

	fmt.Println("\n=== WithValue ===")
	type key int
	const traceKey key = 0
	ctx3 := context.WithValue(context.Background(), traceKey, "trace-123")
	fmt.Println("  traceID:", ctx3.Value(traceKey))
}
