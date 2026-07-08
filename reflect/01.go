package main

import "fmt"

type Book struct {
	Name string
}

// 接口
type Read interface {
	ReadBook()
}

// 接口
type Write interface {
	WriteBook()
}

func (this *Book) ReadBook() {
	fmt.Println("read book")
}

func (this *Book) WriteBook() {
	fmt.Println("read book")
}

func main() {
	// var a string
	// // a -->  <type:"string ", value:"acedl" > go 中每一个变量 都有 这样的对 ( pair ) 分别代表这个变量的type 和 value
	// // 赋值过程中 这个pair对 保持传递  内容不变
	// a = "acedl"

	var b *Book = &Book{}
	var r Read
	r = b

	var w Write
	// r  能直接断言成 w 是以为 b 同时满足这两个接口, 在赋值过程中  pair 对 是直接传递的 不会改变 所以 r 和 w的 pair对 其实是一样的
	w = r.(Write)
	// 既然type 一样又为什么要断言呢？ -- 编译器只能看见 static type 静态类型 表面的类型
	// 表面上 一个是 write 一个是 read 所以需要断言转换
	// 但实际上的type是一眼的
	fmt.Println("r的 类型是", r)
	fmt.Println("w的 类型是", w)

}
