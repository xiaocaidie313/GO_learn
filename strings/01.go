package main

import (
	"fmt"
	"strings"
)

func main() {
	// string 不能改  read-only
	// str := "this is a string"
	// str[0] = 'a' // 无法通过编译
	// fmt.Println(str)
	str2 := "this is a string"
	//读单个字符只能读出 字节值
	fmt.Println(str2[2])
	// 显式类型转换为字节切片
	bytes := []byte(str2)
	fmt.Println(bytes)
	// 显式类型转换为字符串
	fmt.Println(string(bytes))
	// 修改字节切片
	bytes = append(bytes, 96, 97, 98, 99) // 字节切片可改
}

func COPY() {
	str1 := "this is an string"
	str2 := make([]byte, len(str1))
	// copy 要满足两个长度一样
	copy(str2, str1)
	fmt.Println(str2)

	str3 := strings.Clone(str1)

	fmt.Println(str3)

}

func Build() {
	// 高性能的字符串拼接
	builder := strings.Builder{}
	builder.WriteString("this is a string ")
	builder.WriteString("that is a int")
	fmt.Println(builder.String())

}
