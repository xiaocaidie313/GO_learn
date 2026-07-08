package main

import (
	"fmt"
)

// 本征是指针  一套函数规范  c++ 中的多态 template
type Animal interface {
	speak()
}

type Dog struct {
	Name string
	Age  string
}

// speak 方法的 Dog实现
func (this Dog) speak() {
	fmt.Println("我是Dog  Dog")
}

// 当dog 类的方法包含 了 Animal interface 所约束的所方法则 说 dog 就属于 Animal 类  Animal 类是父类
// 注意 必须要完全一样 如果有 返回值不一样 或者 参数不一样 则就不算完成了
type Cat struct {
	Name string
	Age  string
}

// speak 方法的 Cat实现
func (this Cat) speak() {
	fmt.Println("我是Cat  Cat")
}

// 看作是一个接口  里面包括了 这个接口所限制的类型的 所有所需要的函数,  从使用上看 就是这个 接口所具有的功能
func Talk(a Animal) {
	a.speak()
}

// 万能类型 interface {}
// int / float / double 都是 继承这个而来

func myFunc(arg interface{}) {

	// 断言  判断arg是不是 string 类型
	value, ok := arg.(string)
	if ok {
		fmt.Println("我是字符串")
	} else {
		fmt.Println("我不是字符串, 是")
		fmt.Println(value)
	}

}

func main() {

	// d := Dog{Name: "dd", Age: "99"}
	// Talk(d)

	var animal Animal // animal 是父类的指针
	cat := Cat{"cc", "22"}
	animal = &cat
	fmt.Println(animal)
	animal.speak() //
}
