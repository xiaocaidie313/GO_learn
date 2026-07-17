package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// stdout 直接写
	os.Stdout.WriteString("hello stdout!\n")

	// bufio 缓冲写（性能更好）
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("hello bufio!\n")
	writer.Flush()

	// fmt 格式化（有开销，但方便）
	fmt.Println("hello fmt!")
}
