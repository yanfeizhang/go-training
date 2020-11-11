package main

import "fmt"

const (
	x1 = 1
	x2 = "This is a string"
)

const (
	y1 = 100 //100
	y2       //100
	y3       //100
)

const (
	z1 = iota //0
	z2        //1
	z3        //2
)

const (
	_  = 1 << (10 * iota) //1
	KB                    //1024
	MB                    //1048576
	GB                    //1073741824
)

func main() {
	fmt.Println("x1:", x1)
	fmt.Println("x2:", x2)

	fmt.Println("y1:", y1)
	fmt.Println("y2:", y2)
	fmt.Println("y3:", y3)

	fmt.Println("z1:", z1)
	fmt.Println("z2:", z2)
	fmt.Println("z3:", z3)

	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
}
