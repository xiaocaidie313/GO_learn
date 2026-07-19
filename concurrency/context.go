package main

import (
	"context"
	"fmt"
	"time"
)

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
	fmt.Println("=== WithCancel：手动取消 ===")
	ctx1, cancel := context.WithCancel(context.Background())
	go worker(ctx1, "worker-A")
	time.Sleep(700 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== WithTimeout：超时自动取消 ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel2()
	go worker(ctx2, "worker-B")
	time.Sleep(800 * time.Millisecond)

	fmt.Println("\n=== WithValue：传递请求级元数据 ===")
	type key int
	const traceKey key = 0
	ctx3 := context.WithValue(context.Background(), traceKey, "trace-123")
	fmt.Println("  traceID:", ctx3.Value(traceKey))

	fmt.Println("\n=== Context 内部分层（标准库实现思路）===")
	fmt.Println("  Background → WithValue → WithCancel → WithTimeout")
	fmt.Println("  每层只加一种能力，像洋葱一样包起来")
}
