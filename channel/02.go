package main

func main() {
	c := make(chan int, 3)
	// 有缓冲 vs 无缓冲
	// 有缓冲部分还是 是 make(chan int, n) 最大容量是n 当到达最大容量的时候 发生阻塞 直到 管道被消费掉一个 才能继续写入channle中
	// 无缓冲部分是 chan 的 容量是1 写入后一直阻塞  直到被读取 再开放

	// 关闭 channel
	close(c)
	//  select 与 range
}
