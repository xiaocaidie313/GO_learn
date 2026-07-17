# example

## 要点

- 演示**可被其他包 import 的包**结构
- 首字母大写函数/变量 = 导出（public）
- 供 `import/` 目录的点导入示例使用

## 导出函数

| 函数 | 说明 |
|------|------|
| `SayHi()` | 打印问候语 |

## 示例文件

- `01.go` — package example
- `internalTest/01.go` — 普通子包（非 internal 目录）

## 运行

```powershell
go run ./import/   # 通过 import 包间接调用
```
