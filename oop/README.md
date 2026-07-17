# oop

## 要点

- Go 无 class，用 **struct + 方法** 实现 OOP
- **嵌入**字段实现组合式"继承"：`type Human struct { Hero }`
- **首字母大写**导出，小写包内私有
- **interface 隐式实现**：实现所有方法即满足接口
- `interface{}`（或 `any`）可存任意类型，用类型断言取具体类型

## 常用写法

```go
func (h Hero) GetName() { ... }     // 方法
type Animal interface { speak() }   // 接口
Talk(dog)                           // 多态
v, ok := arg.(string)               // 断言
```

## 示例文件

- `01.go` — struct、方法、嵌入
- `interface.go` — interface、多态、断言

## 运行

```powershell
go run ./oop/
```
