# render（响应渲染）

把数据写成 HTTP 响应：JSON / 文件 / 静态资源等。

## 要点

```
c.XXX(状态码, 数据)
→ Gin 序列化并设置 Content-Type
```

前后端分离时最常用 `c.JSON`；`HTML` 较少；文件用 `File` / `Static`。

## 常用方法

| 方法 | Content-Type / 行为 |
|------|---------------------|
| `c.JSON` | `application/json`（默认会转义 HTML 字符） |
| `c.SecureJSON` | JSON 数组前加前缀，防 `<script>` 劫持 |
| `c.PureJSON` | 不转义 HTML 字符的原始 JSON |
| `c.XML` / `YAML` / `TOML` / `ProtoBuf` | 其他格式 |
| `c.File(path)` | 返回本地文件（浏览器可能预览） |
| `c.FileAttachment(path, name)` | 强制下载，自定义文件名 |
| `c.FileFromFS(path, fs)` | 从 `http.FileSystem` 提供 |
| `r.Static` / `StaticFS` / `StaticFile` | 挂静态目录或单文件 |

## 文件

- `01.go` — JSON / SecureJSON / PureJSON / Static / File 示例

## 运行

```powershell
go run ./Gin/render/01.go
# GET /testing
# GET /secure-json
# GET /pure-json
# GET /assets/...   （需有 ./assets 目录时可测）
```
