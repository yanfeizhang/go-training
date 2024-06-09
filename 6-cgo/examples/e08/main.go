package main

//int sum(int a, int b);
//int someCFunc(int n);
import "C"

//export sum
func sum(a, b C.int) C.int {
	return a + b
}

//export someCFunc
func someCFunc(n C.int) C.int {
	return 1
}

func main() {}
