# enum

## 要点

- Go 没有内置 enum，用 **自定义类型 + const + iota** 模拟
- `type Season int` 定义新类型，常量只能取定义过的值
- `iota` 在 `const` 块中从 0 自增，每行 +1
- 也可用 `type WeekDay string` 做字符串枚举

## 常用写法

```go
type Season int
const (
    Spring Season = iota + 1
    Summer
    Autumn
    Winter
)
```

## 示例文件

- `01.go` — Season（iota）与 WeekDay（string）枚举

## 运行

```powershell
go run ./enum/
```
