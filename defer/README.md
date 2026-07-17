# defer

## 要点

- `defer` 在函数 **return 之前** 执行，类似析构
- 多个 defer 按 **后进先出（栈）** 顺序执行
- defer 的**函数参数在 defer 注册时**就求值，不是执行时
- 命名返回值 + defer 可在 return 之间修改最终返回值
- `return` 编译器拆成：赋值 → 执行 defer → 真正返回

## 常用写法

```go
defer fmt.Println("cleanup")
defer func() { ... }()
defer f(x)           // x 在 defer 注册时就算好
func f() (x int) { defer ...; return }  // 命名返回值
```

## 示例文件

- `01.go` — defer 顺序、参数预计算、命名返回值
- `01_test.go` — `TestDefer` 测试

## 运行

```powershell
go run ./defer/
go test ./defer/ -v
```
