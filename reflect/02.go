package main

import (
	"fmt"
	"reflect"
)

// 反射： 通过 ValueOf  TypeOf 函数 反射出 变量的 底层的type 和  value

func refelctArg(arg interface{}) {
	fmt.Println("arg的type是", reflect.ValueOf(arg))
	fmt.Println("arf的类型是", reflect.TypeOf(arg))

}

func main() {
	c := 10
	refelctArg(c)
}
