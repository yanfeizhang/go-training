package main

import (
	"fmt"
	"sync"
)

var (
	sum  = 0
	wg   sync.WaitGroup
	lock sync.Mutex
)

func Add() {
	//如果不加锁会导致两个协程之间覆盖和冲突
	//只有使用锁才能得到符合预期的 20000
	lock.Lock()
	for i := 0; i < 10000; i++ {
		sum = sum + 1
	}
	lock.Unlock()
	wg.Done()
}

func main() {
	wg.Add(2)
	go Add()
	go Add()
	wg.Wait()

	fmt.Println(sum)
}
