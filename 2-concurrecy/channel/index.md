goroutine 之间可以通过 channel 来交换数据

Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

## 无缓存的通道
```go
//创建无缓存的通道
c1 := make(chan int)

//尝试读取，没有数据就会被阻塞
<-c1

//尝试写入，没有人接收也会被阻塞
c1 <- 100
```

## 有缓存的通道
缓冲满载（发送）或变空（接收）之前通信不会阻塞，当满了以后会阻塞掉。

```go
func main() {
	//有缓存的 channel
	c1 := make(chan int, 10)
}
```

## forrange遍历channel
使用 forrange 遍历 channel 很方便，当 channel 被 close 的时候，forrange 会自动退出。

```go
for i := range c {
    fmt.Println(i)
}
```
