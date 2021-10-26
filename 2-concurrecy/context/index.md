当一个计算任务被goroutine承接了之后，由于超时等原因退出的时候，我们希望中止这个goroutine的计算任务。 这个时候可能就需要用到context了。

> 如果不需要为goroutine设置超时，也不需要手工取消执行。只是等待goroutine自己结束的话，只需要WaitGroup即可，不需要context。

# 超时退出
设置context的超时时间
```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
```

定义一个协程函数
```go
func handle(ctx context.Context) {
	ch := make(chan struct{}, 0)
	go func() {
		// 模拟4秒耗时任务
		time.Sleep(time.Second * 2)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
```

在创建goroutine的时候把上面返回的context传入到协程中，并让主协程多等待一会儿。
```go
go handle(ctx)
time.Sleep(time.Second * 5)
```

运行后输出
```go
sub routine timeout
main routine end
```

# context.Background 与 context.TODO
我们查看一下源码
```go
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)
func Background() Context {
	return background
}
func TODO() Context {
	return todo
}
```

从源代码来看，context.Background 和 context.TODO 没有太大的差别。它们只是在使用和语义上稍有不同：

- context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来；
- context.TODO 应该只在不确定应该使用哪种上下文时使用；

在多数情况下，如果当前函数没有上下文作为入参，我们都会使用 context.Background 作为起始的上下文向下传递。

# 取消
除了超时以外，当最上层的 Goroutine 因为某些原因执行失败时，一般也需要将其下面的协程都退出掉。
这时也可以使用context，可以使用context.WithCancel函数，这里就不多举例子了。






