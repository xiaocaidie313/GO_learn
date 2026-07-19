package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Add：计数 +N
	// Done：计数 -1（常用 defer wg.Done()）
	// Wait：计数为 0 前阻塞
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Printf("  worker %d 完成\n", n)
		}(i)
	}

	fmt.Println("主协程等待中...")
	wg.Wait()
	fmt.Println("全部完成，主协程继续")
}
