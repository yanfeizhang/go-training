package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func worker() {
	fmt.Println("Worker is working!")
	wait.Done()
}

func main() {
	wait.Add(1)
	go worker()
	wait.Wait()
}
