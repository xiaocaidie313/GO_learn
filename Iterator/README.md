# Iterator

## 要点

- Go 迭代器：逐个产出元素，不暴露内部结构
- **闭包迭代器**：`func() (T, bool)`，经典写法
- **Go 1.23+**：`iter.Seq[T]` + `for range`，更简洁
- `yield` 不是关键字，只是回调参数名

## 两种写法

```go
// 闭包
next := Fibonacci(10)
for v, ok := next(); ok; v, ok = next() { }

// iter.Seq
for v := range FibonacciSeq(10) { }
```

## 示例文件

- `01.go` — 闭包迭代器 Fibonacci
- `02.go` — `iter.Seq` Fibonacci

## 运行

```powershell
go run ./Iterator/
```
