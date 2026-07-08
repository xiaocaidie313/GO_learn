package main

import (
	_ "init/lib1"      // 导入包 只执行init函数  匿名
	mylib2 "init/lib2" // init 别名
)

func main() {
	// lib1.Lib1Test()
	mylib2.Lib2Test()
}
