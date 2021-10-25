package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {

	fmt.Println("Worker job start...")

	//for range 可以感知到 channle 上的 close
	//当close发生的时候，会退出循环，使用很方便
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Worker job end")
}
func main() {
	//无缓存的 channel
	c1 := make(chan int)
	go worker(c1)

	//循环灌入数据
	for i := 0; i < 100; i++ {
		c1 <- i
	}

	//当结束的时候 close 关闭即可
	close(c1)
	time.Sleep(time.Second)
}
