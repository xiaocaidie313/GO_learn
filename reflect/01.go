package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Name string
}

type Read interface {
	ReadBook()
}
type Write interface {
	WriteBook()
}

func (b *Book) ReadBook()  { fmt.Println("read book:", b.Name) }
func (b *Book) WriteBook() { fmt.Println("write book:", b.Name) }

func main() {
	b := &Book{Name: "Go语言圣经"}
	var r Read = b
	r.ReadBook()

	// 接口断言：Read → Write（b 同时实现了两个接口）
	w := r.(Write)
	w.WriteBook()

	// 反射
	fmt.Println("TypeOf:", reflect.TypeOf(b))
	// 反射基础（见 02.go）
	reflectArg(10)
	reflectArg("hello")

	// Kind：底层种类（见 03 示例）
	var x any = 100
	fmt.Println("TypeOf:", reflect.TypeOf(x), "Kind:", reflect.TypeOf(x).Kind())
}
