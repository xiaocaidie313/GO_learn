# panic

## 要点

- `panic` 会**立即中断**当前函数，已注册的 defer 仍会执行
- panic **之后**的 defer **不会**注册/执行
- 典型顺序：正常代码 → panic 前 defer（LIFO）→ panic 输出
- 生产环境用 `recover()` 捕获 panic（在 defer 中）

## 执行顺序示例

```
C → panic → B(defer) → A(defer) → panic: xxx
```

## 示例文件

- `01.go` — panic 与 defer 执行顺序

## 运行

```powershell
go run ./panic/
```
