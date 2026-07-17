# map

## 要点

- map 是键值对集合，**引用类型**
- 零值 map 为 nil，不能直接写入，需 `make` 或字面量初始化
- `range` 遍历键值对
- `delete(m, key)` 删除键

## 常用方法

| 操作 | 语法 |
|------|------|
| 创建 | `make(map[K]V, cap)` |
| 字面量 | `map[int]string{1: "a"}` |
| 读写 | `m[k] = v`, `v := m[k]` |
| 遍历 | `for k, v := range m` |
| 删除 | `delete(m, k)` |

## 示例文件

- `01.go` — 创建、遍历、delete

## 运行

```powershell
go run ./map/
```
