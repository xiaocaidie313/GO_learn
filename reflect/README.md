# reflect

## 要点

- 每个变量是 **(type, value)** 对
- interface 有**静态类型**（编译器看到的）和**动态类型**（实际存的）
- 接口断言 `r.(Write)` 在编译类型不同时需要显式转换
- 反射：`reflect.TypeOf(v)` 取类型，`reflect.ValueOf(v)` 取值

## 常用方法

| 方法 | 说明 |
|------|------|
| `reflect.TypeOf(x)` | 运行时类型 |
| `reflect.ValueOf(x)` | 运行时值 |
| `v.(T)` | 类型断言 |
| `v, ok := x.(T)` | 安全断言 |

## 示例文件

- `01.go` — 接口静态/动态类型、断言
- `02.go` — TypeOf / ValueOf 基础

## 运行

```powershell
go run ./reflect/
```
