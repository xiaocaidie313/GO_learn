package main

import "fmt"

type Dog struct{ Name string }
type Cat struct{ Name string }

func (d Dog) speak() { fmt.Println("汪汪", d.Name) }
func (c Cat) speak() { fmt.Println("喵喵", c.Name) }

type Animal interface {
	speak()
}

func Talk(a Animal) {
	a.speak()
}

func myFunc(arg interface{}) {
	if v, ok := arg.(string); ok {
		fmt.Println("是字符串:", v)
	} else {
		fmt.Println("不是字符串")
	}
}
