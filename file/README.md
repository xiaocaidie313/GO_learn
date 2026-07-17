# file

## 要点

- 文件操作在 `os` 包中
- 打开文件后要 `defer file.Close()`
- 读文件常用 `file.Read(buffer)`，注意 buffer 扩容
- 遍历目录用 `os.ReadDir`

## 常用方法

| 方法 | 说明 |
|------|------|
| `os.Open(path)` | 打开文件 |
| `os.IsNotExist(err)` | 判断文件不存在 |
| `(*os.File).Read(buf)` | 读取字节 |
| `file.Close()` | 关闭文件 |
| `os.ReadDir(path)` | 读取目录条目 |

## 示例文件

- `01.go` — 打开文件、ReadFile、ReadDir

## 运行

```powershell
go run ./file/
```
