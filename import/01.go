package importTest

// 包是基本单位  例如我有多个文件 在不同文件位置  但是都在顶部 package xxx 声明是同一个包

// 1. 包的导入
// 无法实现 循环导入
import (
	. "demo/example" // 全部导入  不使用  包名.func()
	_ "fmt"          // 匿名导入包  不是用包内函数 只是使用包的Init函数
	t "time"         // 取别名
	// 	 internalTest "demo/example/internal"  // 无法导入 internalTest 包 因为 internalTest 包是 internal 包 内部包 外部无法访问
	// 只能被 demo/example 包 内部访问 父目录下的包访问
)

// 2 为什么无法实现循环导入？  因为每个包导入的时候都有Init函数 如果 循环导入了 不知道先调用哪个Init函数
// A 中 import B， B中import A 这样就会形成循环导入 导致编译错误  ---->应该提取公共方法 到第三方包中的

func main() {

	t.Now()
	SayHi()
}
