package main

import "fmt"

func SliceAppend1() {
	var t []int
	for i := 0; i < 1000000; i++ {
		t = append(t, i)
	}
}

func SliceAppend2() {
	t := make([]int, 0, 1000000)
	for i := 0; i < 1000000; i++ {
		t = append(t, i)
	}
}

func main() {
	fmt.Println("hello world")
}
