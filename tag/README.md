# tag

## 要点

- struct 字段后的 **tag** 是元数据字符串，供反射/序列化使用
- JSON 序列化用 `encoding/json`，tag 格式：`json:"fieldName"`
- tag **不能有空格**：`` `json:"name"` `` ✅，`` `json : "name"` `` ❌
- `json.Marshal` → `[]byte`，`json.Unmarshal` → 结构体

## 常用方法

| 方法 | 说明 |
|------|------|
| `json.Marshal(v)` | 结构体 → JSON 字节 |
| `json.Unmarshal(data, &v)` | JSON → 结构体 |
| `` `json:"key"` `` | 指定 JSON 字段名 |

## 示例文件

- `01.go` — struct tag + JSON 序列化/反序列化

## 运行

```powershell
go run ./tag/
```
