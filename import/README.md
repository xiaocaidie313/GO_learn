# import

## 要点

- 包是 Go 的基本编译单位，同包多文件共享 `package xxx`
- **不能循环导入**（init 顺序无法确定）
- 目录名 `internal` 有特殊含义：外部包不能导入

## 四种导入方式

| 写法 | 含义 |
|------|------|
| `"demo/example"` | 正常导入，用 `example.SayHi()` |
| `. "demo/example"` | 点导入，直接 `SayHi()` |
| `_ "fmt"` | 匿名导入，只用包的 init |
| `t "time"` | 别名导入，用 `t.Now()` |

## 示例文件

- `01.go` — 四种 import 写法演示

## 运行

```powershell
go run ./import/
```
