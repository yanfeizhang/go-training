package main

func main() {
	n := 1
	_ = func1(n)
	_ = func2(n)
}

func func1(n int) int {
	_ = make([]byte, 200)
	return n
}

func func2(n int) int {
	_ = make([]byte, 20)
	return n
}
