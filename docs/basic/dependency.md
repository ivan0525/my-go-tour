# 依赖管理

## 1. `go mod`命令

常用的`go mod`命令

```bash
# 下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod download 

# 编辑go.mod文件
go mod edit

# 打印模块依赖图
go mod graph

# 初始化当前文件夹，创建go.mod
go mod init

# 增加缺少的module，删除无用的module
go mod tidy

# 将依赖复制到vendor下
go mod vendor

# 校验依赖
go mod verify

# 解释为什么需要依赖
go mod why
```