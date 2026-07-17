package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello"
	bytes := []byte(str)
	bytes[0] = 'H' // 通过 []byte 间接修改
	fmt.Println(string(bytes)) // Hello

	// Builder 高效拼接
	var b strings.Builder
	b.WriteString("Go ")
	b.WriteString("语言")
	fmt.Println(b.String())

	// Clone 深拷贝
	s1 := "immutable"
	s2 := strings.Clone(s1)
	fmt.Println(s1, s2)
}
