# var

## 要点

- 四种变量声明：`var x int`、`var x = 1`、`var x, y = 1, "a"`、`x := 1`
- `:=` **只能在函数内**使用
- **常量**用 `const`，只读
- **iota** 是 const 块内的自增计数器，每行 +1

## iota 规则

- 遇到新 `const (` 块时重置为 0
- 同一行多个常量共享当前 iota
- 可省略重复表达式，继承上一行

## 示例文件

- `test_01.go` — var、:=、const、iota
- `swap/01.go` — 并行赋值交换

## 运行

```powershell
go run ./var/test_01.go
go run ./var/swap/
```
