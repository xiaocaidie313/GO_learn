package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== 无缓冲 channel：发送和接收必须同时就绪 ===")
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  发送 42")
		ch <- 42
	}()
	v := <-ch
	fmt.Println("  收到:", v)
	wg.Wait()

	fmt.Println("\n=== 有缓冲 channel：缓冲区未满时不阻塞 ===")
	bufCh := make(chan int, 2)
	bufCh <- 1
	bufCh <- 2
	fmt.Println("  读出:", <-bufCh, <-bufCh)

	fmt.Println("\n=== 关闭 channel：range 读至结束 ===")
	done := make(chan int, 3)
	done <- 1
	done <- 2
	done <- 3
	close(done)
	for v := range done {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
