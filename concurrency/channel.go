package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建了一个无缓冲管道
	// 无缓冲部分是 chan 的 容量是0  写入后一直阻塞  直到被读取 再开放
	fmt.Println("=== 无缓冲 channel ===")
	intCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  发送 1")
		intCh <- 1 // 如果协程先到 c <- 1 则会在协程发生阻塞 直到 main 的 <-c 读出
	}()
	// 同理如果 <-c 先到 则会阻塞  等待 c<- 的写入
	ints, ok := <-intCh
	fmt.Println("  读取:", ints, ok)
	wg.Wait()

	// 有缓冲 vs 无缓冲
	// 有缓冲部分是 make(chan int, n) 最大容量是n 当到达最大容量的时候 发生阻塞 直到 管道被消费掉一个 才能继续写入channel中
	fmt.Println("\n=== 有缓冲 channel ===")
	intCh2 := make(chan int, 1)
	intCh2 <- 1
	ints2, ok2 := <-intCh2
	fmt.Println("  读出:", ints2, ok2)

	// 访问空的有缓冲的管道 —— 缓冲区为空，阻塞等待其他协程写入数据
	// intCh3 := make(chan int, 1)
	// ints3, ok3 := <-intCh3  // 会阻塞，示例中注释掉

	// 写入满管道 —— 满了，阻塞等待其他协程来读取数据
	// intCh4 := make(chan int, 1)
	// intCh4 <- 1
	// intCh4 <- 2  // 会阻塞，示例中注释掉

	// 关闭 channel
	fmt.Println("\n=== close + range ===")
	done := make(chan int, 3)
	done <- 1
	done <- 2
	done <- 3
	close(done)
	// select 与 range
	for v := range done {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 管道为nil
	// var intCh5 chan int  // 只是声明了变量 但是没有分配空间，初始化
	// intCh5 <- 1          // 会永久阻塞
	// close(nilCh)         // panic
}
