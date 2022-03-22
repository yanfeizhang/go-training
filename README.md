# go-training
目前市面上各种golang的学习资料实在不少，不过学习过一段时间以后发现有两个问题。一是每一个文章、站点都是介绍一个方面，缺乏一个整体上的视角。二是学习完了之后不练习的话，基本一周后就忘的差不多了。

所以我把我个人的学习、练习过程以此项目整理发布了出来。
- 用树形的方式进行尽量合理地组织，便于整体上把握
- 每一个知识点尽量配置go源代码，便于动手练习

该项目目前只更新了基础部分，日后会继续更新和完善。

![avatar](imgs/golang-small.png)

# 开发环境
- 安装
- 配置GOROOT、GOPATH
- GoLand使用以及其快捷键

# Golang知识树
- 语言基础：
    - [关键字与预定义标识符](1-basic/keyword/index.md)  
    - [常量的声明与赋值](1-basic/const/index.md) &nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/const/main.go)
    - [变量的声明、初始化、生命期](1-basic/variable/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/variable/main.go)
    - [流程控制：if/else、for、switch](1-basic/flow/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/flow/main.go)  
 
- 数据类型
    - [数据类型之整形、布尔、字符串](1-basic/int-bool-string/index.md)
    - [指针数据类型](1-basic/pointer/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code]((1-basic/pointer/main.go))                                                               
    - [数据类型之数组、切片](1-basic/array/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/array/main.go)
    - [数据类型之Map](1-basic/map/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/map/main.go)
    - [数据类型之结构体](1-basic/struct/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/struct/main.go)
    - [数据类型之接口](1-basic/interface/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/interface/main.go)
    - [类型转换](1-basic/convert/index.md) &nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/convert/main.go)
    - [类型断言](1-basic/assert/index.md) &nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/assert/main.go)
    - [反射](1-basic/reflect/index.md) &nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/reflect/main.go)
    
- 代码组织 
    - [函数、匿名函数、闭包](1-basic/function/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/function/main.go)
    - [方法](1-basic/method/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/method/main.go)
    - [包](1-basic/package/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/package/main.go)  
    - [defer、panic和recover](1-basic/defer/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/defer/main.go) 
    - [go module](1-basic/module/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](1-basic/module/main.go)  

- 并发操作
    - [goroutine](2-concurrecy/goroutine/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](2-concurrecy/goroutine/main.go)
    - [channel](2-concurrecy/channel/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code（无缓存）](2-concurrecy/channel/buffer/main.go)&nbsp;&nbsp;[->code（带缓存）](2-concurrecy/channel/buffer/main.go)&nbsp;&nbsp;[->code（遍历）](2-concurrecy/channel/buffer/main.go)
    - [select](2-concurrecy/select/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](2-concurrecy/select/main.go);
    - [sync](2-concurrecy/sync/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code(waitgroup)](2-concurrecy/sync/waitgroup/main.go)&nbsp;&nbsp;[->code(mutex)](2-concurrecy/sync/mutex/main.go)&nbsp;&nbsp;[->code(rwmutex)](2-concurrecy/sync/rwmutex/main.go);
    - context
        - [context源码浅析](2-concurrecy/context/mycontext/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](2-concurrecy/context/mycontext/main.go)
        - [timeout简单demo](2-concurrecy/context/timeout/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](2-concurrecy/context/timeout/main.go) 
        - [value简单demo](2-concurrecy/context/value/index.md)&nbsp;&nbsp;&nbsp;&nbsp;[->code](2-concurrecy/context/value/main.go) 
- 网络操作
    - [官方net包]() 网络库，支持 Socket、HTTP、邮件、RPC、SMTP 等  [->code](1-basic/function/main.go)
    - [gnet包]()
    - [evio]()
    - [netpoll]()
- 编码解码
    - encoding: 常见算法如 JSON、XML、Base64 等
    - json 
        - [encoding/json](4-codec/json/json/index.md)   [->code](4-codec/json/json/main.go)
        - json-iterator
        - tidwall/gjson  
        - easyjson
    - protobuf        
- 单元测试
    - [基础单元测试](3-unittest/basic/main_test.go)           
    - [testify单元测试](3-unittest/testify/main_test.go)    
    - [gomonkey单元测试](3-unittest/gomonkey/main_test.go)    
- 标准库用法
    - bytes: 字节操作
    - context
    - database: 数据库驱动和接口,包括mysql
    - flag：命令行解析
    - fmt: 格式化操作
    - io：实现 I/O 原始访问接口及访问封装
    - mysql: mysql数据库操作
    - math: 数学库
    - os: 操作系统平台不依赖平台操作封装
    - sort: 排序接口
    - strings: 字符串转换、解析及实用函数
    - time: 时间接口
- golang工具链
    - go get命令: 一键获取代码、编译并安装
    - go install命令: 编译并安装
    - go module: 依赖管理
    - go fmt命令: 格式化代码文件
    - go imports: 
    - go build: 编译
    - go run: 编译并运行
    - go test: 单元测试
    - go pprof命令: 性能分析
    - 数据竞争
    - 代码覆盖率
    
# Golang进阶
- 编译原理
    - [golang之编译原理]
    - [数组编译过程理解]
    - [为什么数组常量下标编译时报错，而变量下标运行时panic]
    - [对切片append超过其容量时会发生什么]
    - [了解字符串拼接、转换的开销]
    - [golang函数调用和c函数开销对比]
    - [为什么make能同时支持map、slice和chan]
    - [for range循环解惑]
    - [GPM之阻塞的系统调用实勘]
    - [Golang协程上的网络IO]
- 主流第三方库用法
    - gorm
    - cobra
    - redigo
    - go-cache
    - viper
    - gRPC   
    - 日志库 
- 并发
    - goroutine
    - channel
    - select
    - sync包：WaitGroup、互斥锁、读写互斥锁、Once、并发安全Map           
- 工程化与部署
    - 编码规范
    - 热升级   
    - 项目布局
    - Make file
    - Docker file 
    - 优雅地调试容器
    - CI && CD
# 一个web项目框架

# 经典问题集锦

|--|--|
|--|--|

# 参考学习资料
- [https://golang.org/doc/code.html](https://golang.org/doc/code.html) 
- [https://golang.org/doc/effective_go.html](https://golang.org/doc/effective_go.html) 
- [http://c.biancheng.net/golang/intro/](http://c.biancheng.net/golang/intro/)
- [awesome-go项目中文版](https://github.com/yinggaozhen/awesome-go-cn)  精选了一系列很棒的Go框架、库和软件。灵感来自于awesome-python
- [Uber 内部在 github 开源的 的 Go 编程规范](https://github.com/uber-go/guide)
- [Dave Cheney 写的 《Go语言最佳实战》 ](https://dave.cheney.net/practical-go/presentations/qcon-china.html)
- [Go语言设计与实现(电子书)](https://draveness.me/golang/)  侧重于内部实现分析
- [begoo官方文档](https://beego.me/)
- [如何优雅地写出Go代码](https://draveness.me/golang-101/)


# 硬广个人技术公众号
![avatar](imgs/wechat.png)