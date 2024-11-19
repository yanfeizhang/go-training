package main

import (
	"fmt"
	"math/rand"
)

type ItemA struct {
	f1  int8
	f2  int
	f3  int8
	f4  int
	f5  int8
	f6  int
	f7  int8
	f8  int
	f9  int8
	f10 int
	f11 int8
	f12 int
	f13 int8
	f14 int
}

func someValueFunc(a ItemA) int {
	return a.f8
}

func somePointerFunc(a *ItemA) int {
	return a.f8
}

func main() {

	fmt.Println(rand.Intn(100))
}
