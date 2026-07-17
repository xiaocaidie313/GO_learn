package main

import (
	"fmt"
	"reflect"
)

type book int

// .kind()返回底层大类

func main() {

	var a book
	// a = 100
	_ = a

	fmt.Println("a的类型是", reflect.TypeOf(a))
	fmt.Println("a的类型名称是  reflect.TypeOf(a).Name()==", reflect.TypeOf(a).Name()) // book

	fmt.Println("a的类型是.kind()", reflect.TypeOf(a).Kind())

}
