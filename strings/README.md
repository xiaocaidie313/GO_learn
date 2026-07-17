# strings

## 要点

- **string 不可变**（read-only），不能直接改 `str[0]`
- 通过 `[]byte(str)` 转换后可修改字节，再转回 `string`
- 大量拼接用 `strings.Builder`，比 `+` 高效
- `strings.Clone(s)` 深拷贝字符串

## 常用方法

| 方法 | 说明 |
|------|------|
| `[]byte(s)` / `string(b)` | string ↔ []byte |
| `strings.Builder` | 高效拼接 |
| `builder.WriteString(s)` | 写入 Builder |
| `builder.String()` | 得到最终 string |
| `strings.Clone(s)` | 拷贝字符串 |

## 示例文件

- `01.go` — 不可变、转换、Builder、Clone

## 运行

```powershell
go run ./strings/
```
