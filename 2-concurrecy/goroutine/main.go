package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello")
}

func main() {
	go hello()
	time.Sleep(time.Hour)
}
