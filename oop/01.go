package main

import (
	"fmt"
)

// 首字母大写的 都是外部能访问的  首字母小写的都是只能在内部访问的  是私有的部分

type Hero struct {
	Name  string
	Ad    string
	level string
}

// 定义子类
type Human struct {
	Hero // 父类
	leg  int
}

// 类似于js对象中的 handler 方法
// 这里的 " this " 只是用于 这个方法内部的暂时的使用变量 其他时候不影响
func (this Hero) GetName() {
	// this 是一个变量 他的类型是 Hero !
	fmt.Println("Name = ", this.Name)
}

// 就是常规的对象方法
func (this Hero) SetName() {

}

func (A Hero) SetAd() {

}

func main() {

	hero := Hero{Name: "张飞", Ad: "2222", level: "999"}
	fmt.Print(hero)
}
