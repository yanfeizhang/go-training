package main

import (
	"fmt"
	"sync"
)

var (
	sum  int = 0
	wg   sync.WaitGroup
	lock sync.RWMutex
)

func write() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		sum = sum + 1
		lock.Unlock()
	}
	wg.Done()
}

func read() {
	for i := 0; i < 10000; i++ {
		lock.RLock()
		fmt.Println(sum)
		lock.RUnlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go write()
	go read()
	wg.Wait()

	fmt.Println("Job finish!")
}
