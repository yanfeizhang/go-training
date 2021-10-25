package main

func main() {
	//有缓存的 channel
	c1 := make(chan int, 10)

	//在添加满通道之前，当前协程都是加完后立即返回
	for i := 0; i < 10; i++ {
		c1 <- i
	}

	//当发送的数据超过通道大小的时候，主协程会阻塞
	//如果没有协程消耗，会导致死锁报错
	c1 <- 100
}
