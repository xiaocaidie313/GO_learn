package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// once
// mutex rwmutex
// pool

/*
在使用 sync.Pool 时需要注意几个点：
临时对象：sync.Pool 只适合存放临时对象，池中的对象可能会在没有任何通知的情况下被 GC 移除，
所以并不建议将网络链接，数据库连接这类存入 sync.Pool 中。
不可预知：sync.Pool 在申请对象时，无法预知这个对象是新创建的还是复用的，也无法知晓池中有几个对象
并发安全：官方保证 sync.Pool 一定是并发安全，但并不保证用于创建对象的 New 函数就一定是并发安全的，New 函数是由使用者传入的，
所以 New 函数的并发安全性要由使用者自己来维护，这也是为什么上例中对象计数要用到原子值的原因。
*/

func demoMutex() {
	fmt.Println("=== Mutex ===")
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
	fmt.Println("\n=== Once ===")
	var once sync.Once
	for i := 0; i < 3; i++ {
		once.Do(func() {
			fmt.Println("  初始化（只会打印一次）")
		})
	}
}

func demoPool() {
	fmt.Println("\n=== Pool ===")
	var count int64
	pool := sync.Pool{
		New: func() any {
			atomic.AddInt64(&count, 1) // 对象计数用原子值
			return make([]byte, 1024)
		},
	}
	buf := pool.Get().([]byte)
	pool.Put(buf)
	buf2 := pool.Get().([]byte)
	pool.Put(buf2)
	fmt.Println("  New 被调用次数:", atomic.LoadInt64(&count))
}

func main() {
	demoMutex()
	demoOnce()
	demoPool()
}
