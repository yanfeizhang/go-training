package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	for i := 0; i < 10000000; i++ {
		ch <- i
	}
}

func stoper(ch chan int) {
	time.Sleep(time.Second)
	ch <- 0
	time.Sleep(time.Second)
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go worker(ch1)
	go worker(ch2)
	go stoper(ch3)

	for {
		select {
		case i := <-ch1:
			fmt.Println("Worker1 job done", i)
		case j := <-ch2:
			fmt.Println("Worker2 job done", j)
		case _, ok := <-ch3:
			if ok {
				fmt.Println("Job continue")
			} else {
				fmt.Println("Kill all job")
				return
			}
		}
	}
}
