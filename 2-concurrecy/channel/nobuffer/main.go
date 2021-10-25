package main

import (
	"fmt"
	"time"
)

func worker(c chan int){
	fmt.Println("Worker job start...")

	//故意延迟灌入数据
	time.Sleep(time.Second)
	c <- 100
	fmt.Println("Worker job end")
}
func main(){
	//无缓存的 channel
	c1 := make(chan int)
	go worker(c1)

	//主协程阻塞在这里，等待灌入数据
	fmt.Println(<-c1)
}
