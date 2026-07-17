package main

// Fibonacci 闭包迭代器：每次调用 next() 返回 (值, 是否还有下一个)
func Fibonacci(n int) func() (int, bool) {
	a, b, c := 1, 1, 2
	i := 0
	return func() (int, bool) {
		if i >= n {
			return 0, false
		} else if i < 2 {
			f := i
			i++
			return f, true
		}

		a, b = b, c
		c = a + b
		i++

		return a, true
	}
}
