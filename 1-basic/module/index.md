
# init
在初始化创建一个项目的时候
```go  
go mod init
```

这时候会创建一个go.mod文件，其内容如下。
```go 
module module

go 1.14
```
这时再执行go build 等命令就会下载依赖包，并把依赖信息添加到 go.mod 文件中，同时把依赖版本哈希信息存到 go.sum 文件中
例如当我们在main.go使用了pflag之后，
```go  
go build .
go: finding module for package github.com/spf13/pflag
go: found github.com/spf13/pflag in github.com/spf13/pflag v1.0.5
```

# tidy
当工程运行过一段时间以后，难免会导致go.mod中存在一些已经不用了的依赖。这时候可以通过tidy来清除
```go  
go mod tidy
```
