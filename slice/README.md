# slice

## 要点

- **数组**长度固定，是值类型；**切片**动态长度，是引用类型
- 切片底层：`{ptr, len, cap}`
- `array[a:b]` 切片与数组**共享底层数组**
- `append` 可能触发扩容，返回新 slice 头
- `copy(dst, src)` 复制元素（取较短长度）

## 常用方法

| 操作 | 语法 |
|------|------|
| 创建 | `make([]T, len, cap)` |
| 追加 | `s = append(s, v...)` |
| 复制 | `copy(dst, src)` |
| 克隆 | `slices.Clone(s)` |
| 长度/容量 | `len(s)`, `cap(s)` |
| 删除 | `s = append(s[:i], s[i+1:]...)` |

## 示例文件

- `01.go` — 数组/切片、append、insert、delete、cap

## 运行

```powershell
go run ./slice/
```
