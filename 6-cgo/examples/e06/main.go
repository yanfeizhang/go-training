package main

/*
#include <errno.h>

static int someCFunc(int a) {
    return 1;
}
*/
import "C"
import (
	"fmt"
	"time"
)

func someGoFunc(a int) int {
	return 1
}

func main() {
	var testCount int64 = 100000000
	var i int64

	// Go func test
	start := time.Now()
	for i = 0; i < testCount; i++ {
		someGoFunc(5)
	}
	end := time.Now()
	fmt.Println("Go 函数调用开销:", float64(end.Sub(start).Nanoseconds())/float64(testCount))

	// CGO func test
	start = time.Now()
	for i = 0; i < testCount; i++ {
		C.someCFunc(5)
	}
	end = time.Now()
	fmt.Println("CGO 函数调用开销:", float64(end.Sub(start).Nanoseconds())/float64(testCount))
}
