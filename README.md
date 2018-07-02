## 构建
```
go build -o server main.go
```

## 安装问题
1. 提示

> [ERROR]	Failed to set version on golang.org/x/sys/unix to : Cannot detect VCS
  [ERROR]	Failed to set references: Cannot detect VCS

是下载golang.org/x/sys/unix失败导致的，需执行如下命令

```
glide mirror set https://golang.org/x/sys/unix https://github.com/golang/sys
```