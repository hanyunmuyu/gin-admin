[使用alpine 构建 golang 运行容器](https://www.cnblogs.com/lovezbs/p/13199121.html)

Mac 下编译 Linux 和 Windows 64位可执行程序

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

```

