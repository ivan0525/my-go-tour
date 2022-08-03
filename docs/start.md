# 起步

## 设置`GOPROXY`

由于默认的`https://proxy.golang.org`在国内可能无法访问，所以需设置国内的模块代理。

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```