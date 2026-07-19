package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func demoMutex() {
	fmt.Println("=== Mutex 保护共享变量 ===")
	var (
		mu    sync.Mutex
		count int
		wg    sync.WaitGroup
	)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("  count =", count)
}

func demoOnce() {
	fmt.Println("\n=== Once 只执行一次 ===")
	var once sync.Once
	for i := 0; i < 3; i++ {
		once.Do(func() {
			fmt.Println("  初始化（只会打印一次）")
		})
	}
}

func demoPool() {
	fmt.Println("\n=== Pool 临时对象复用 ===")
	var count int64
	pool := sync.Pool{
		New: func() any {
			atomic.AddInt64(&count, 1)
			return make([]byte, 1024)
		},
	}
	buf := pool.Get().([]byte)
	pool.Put(buf)
	buf2 := pool.Get().([]byte) // 可能复用，也可能新建
	pool.Put(buf2)
	fmt.Println("  New 被调用次数:", atomic.LoadInt64(&count), "（≤2，可能复用）")
	fmt.Println("  注意：Pool 中对象可能随时被 GC 清除，不适合放连接等资源")
}

func main() {
	demoMutex()
	demoOnce()
	demoPool()
}
