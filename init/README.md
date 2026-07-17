# init

## 要点

- 每个包可有 `init()`，在 `main()` 之前自动执行
- **import 包时会执行被导入包的 init**
- `_ "init/lib1"` 匿名导入：不调用函数，只触发 lib1 的 init
- 多包 init 顺序：依赖包先 init，同层按 import 顺序

## 常用写法

```go
import _ "init/lib1"           // 只执行 init
import mylib2 "init/lib2"      // 别名 + 可调用导出函数
```

## 示例文件

- `main.go` — 触发 lib1/lib2 的 init
- `lib1/lib1.go` — 导出函数 + init
- `lib2/lib2.go` — 导出函数 + init

## 运行

```powershell
cd init
go run main.go
```
