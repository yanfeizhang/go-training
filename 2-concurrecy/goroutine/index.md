
在 Go 语言中，每一个并发的执行单元叫作一个 goroutine，goroutine 的概念类似于线程，属于用户态的线程, 是由 Go 的运行时（runtime）调度和管理的。

每一个 OS 线程都有一个固定大小的内存块(一般会是2MB)来做栈，这个栈会用来存储当前正在被调用或挂起(指在调用其它函数时)的函数的内部变量。相反，一个 goroutine 会以一个很小的栈开始其生命周期，一般只需要 2KB。一个 goroutine 的栈，和操作系统线程一样，会保存其活跃或挂起的函数调用的本地变量，但是和 OS 线程不太一样的是一个 goroutine 的栈大小并不是固定的；栈的大小会根据需要动态地伸缩。而 goroutine 的栈的最大值有 1GB，比传统的固定大小的线程栈要大得多，尽管一般情况下，大多 goroutine 都不需要这么大的栈。

## 使用

Go 语言中使用 goroutine 非常简单，只需要在调用函数或方法的时候在前面加上 go 关键字，就可以为一个函数创建一个 goroutine。当一个程序启动时，其 main() 函数即在一个单独的 goroutine 中运行，我们叫它 main goroutine，尽管它并没有通过 go 来启动。

不过由于主线程有可能先于 hello 协程退出，先用最简单的办法，sleep 一秒等待输出。

```go
func hello(){
	fmt.Println("Hello")
}

func main(){
	go hello()
	time.Sleep(1)
}
```

## 调度
GPM 是 Go 语言运行时（runtime）层面的实现，是 go 语言自己实现的一套调度系统。区别于操作系统调度 OS 线程。

- G: Gourtines, 每个 Goroutine 对应一个 G 结构体，G 保存 Goroutine 的运行堆栈，即并发任务状态。G 并非执行体，每个 G 需要绑定到P才能被调度执行。
- P: Processors, 对 G 来说，P 相当于 CPU 核，G 只有绑定到 P(在P的 local runq 中)才能被调度。对M来说，P 提供了相关的执行环境(Context)，如内存分配状态(mcache)，任务队列(G)等，P 会对自己管理的 goroutine 队列做一些调度（比如把占用 CPU 时间较长的 goroutine 暂停、运行后续的 goroutine 等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。
- M（machine）是 Go 运行时（runtime）对操作系统内核线程的抽象， M 与内核线程一般是一一映射的关系， 一个 groutine 最终是要放到M上执行的； P 与 M 一般也是一一对应的。他们关系是： P 管理着一组 G 挂载在 M 上运行。当一个 G 长久阻塞在一个 ，runtime 会新建一个 M，阻塞 G 所在的 P 会把其他的 G 挂载在新建的 M 上。当旧的 G 阻塞完成或者认为其已经死掉时 回收旧的 M。

P 的个数是通过 runtime.GOMAXPROCS 设定（最大256），Go 1.5版本之后默认为物理线程数。
M 的个数是不定的，由 Go Runtime 调整，默认最大限制为 10000 个。