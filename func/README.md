# func

## 要点

- Go 函数可返回**多个值**
- **匿名返回**：`(int, int)` 直接 return 字面量
- **命名返回**：`(res1 int, res2 string)` 可裸 `return`
- 命名返回值在函数开头就初始化

## 常用写法

```go
func add(a, b int) (int, int) { return a + b, a - b }
func foo(a, b int) (sum int, msg string) { sum = a+b; return }
```

## 示例文件

- `01.go` — 多返回值、命名返回对比

## 运行

```powershell
go run ./func/
```
