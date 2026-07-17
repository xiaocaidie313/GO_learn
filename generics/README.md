# generics

## 要点

- Go 1.18+ 支持泛型：`func Do[T Constraint](n T) T`
- **类型约束**用 interface 定义，如 `~int | ~int64`
- `~T` 表示**底层类型**是 T 的类型也满足约束
- 自定义类型 `type TinyInt int8` 不等于 `int8`，需约束里写 `~int8`

## 常用写法

```go
type Number interface { ~int | ~int64 }
func Max[T Number](a, b T) T { ... }
```

## 示例文件

- `01.go` — 类型集约束与 `TinyInt` 编译限制

## 运行

```powershell
go run ./generics/
```
