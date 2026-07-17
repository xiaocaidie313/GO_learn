package main

import (
	"fmt"
	"reflect"
)

func reflectArg(arg interface{}) {
	fmt.Println("TypeOf:", reflect.TypeOf(arg))
	fmt.Println("ValueOf:", reflect.ValueOf(arg))
}
