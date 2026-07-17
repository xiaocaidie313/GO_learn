# inout

## 要点

- **stdin / stdout / stderr** 都是文件描述符
- `fmt.Println` 有格式化开销，高性能场景用 `bufio`
- 读 stdin 多种方式：`Read`、`Scanln`、`bufio.Reader`、`Scanner`

## 常用方法

| 方法 | 说明 |
|------|------|
| `os.Stdout.WriteString(s)` | 直接写 stdout |
| `bufio.NewWriter(os.Stdout)` | 缓冲写 |
| `fmt.Scanln(&a, &b)` | 读一行到变量 |
| `bufio.NewReader(os.Stdin).ReadString('\n')` | 读到换行 |
| `bufio.NewScanner(os.Stdin)` | 逐行扫描 |

## 示例文件

- `01.go` — stdout 写入、bufio、stdin 读取

## 运行

```powershell
go run ./inout/
```
