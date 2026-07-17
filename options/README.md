# options

## 要点

- **Functional Options 模式**：用函数选项替代超长构造函数参数列表
- `type Option func(*Config)` 定义选项函数
- `NewXxx(opts ...Option)` 遍历应用每个选项
- 扩展性好，新增字段只需加 `WithXxx` 函数

## 常用写法

```go
type PersonOption func(*Person)
func WithName(n string) PersonOption { return func(p *Person) { p.Name = n } }
p := NewPerson(WithName("张三"), WithAge(20))
```

## 示例文件

- `01.go` — Person + WithName/WithAge 等选项

## 运行

```powershell
go run ./options/
```
