# swap

## 要点

- Go 支持**并行赋值**：`a, b = b, a` 直接交换，无需临时变量
- `switch` 可替代 if-else 链做分支

## 常用写法

```go
a, b = b, a
switch mode {
case "max": ...
case "min": ...
}
```

## 示例文件

- `01.go` — 交换变量、Compare max/min

## 运行

```powershell
go run ./var/swap/
```
