package main

func main() {
	a, b := 1, 2
	_ = add1(a, b)
	_ = add2(a, b)
	_ = add3(a, b)
}

func add1(x, y int) int {
	_ = make([]byte, 20)
	return x + y
}

func add2(x, y int) int {
	_ = make([]byte, 200)
	return x + y
}

func add3(x, y int) int {
	_ = make([]byte, 5000)
	return x + y
}
