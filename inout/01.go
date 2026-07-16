package inout

import (
	"bufio"
	"fmt"
	"os"
)

// 标准输入输出
// stout本身是一个文件 所以可以写入文件
func main() {
	os.Stdout.WriteString("hello world!")

	// 将内容写入到标准错误中  不建议使用
	print("hello world@@@!\n")
	println("hello world$$$$")

	// 对性能要求场景 不用这个 因为需要格式化 性能损耗
	fmt.Println("___***")

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	writer.WriteString("hello world!")

}

func useBufio() {
	// 使用bufio 写入 性能更高 因为不需要格式化  写入到内存中
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	writer.WriteString("hello world! useBufio")
}

func Read() {
	var buf [1024]byte
	n, _ := os.Stdin.Read(buf[:])
	os.Stdout.Write(buf[:n]) // 将读取到的内容写入到标准输出中

	var a, b int
	fmt.Scanln(&a, &b) // 扫描缓存区的  就是终端上的  然后根据地址 赋值给a和b
	fmt.Printf("%d + %d = %d\n", a, b, a+b)

	//reader
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	fmt.Println("read", line)

	//scanner

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			break
		}
		fmt.Println("scan", line)
	}
}

/*

// 扫描从os.Stdin读入的文本，根据空格分隔，换行也被当作空格
func Scan(a ...any) (n int, err error)

// 与Scan类似，但是遇到换行停止扫描
func Scanln(a ...any) (n int, err error)

// 根据格式化的字符串扫描
func Scanf(format string, a ...any) (n int, err error)


*/
