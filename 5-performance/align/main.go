package main

import (
	"fmt"
	"unsafe"
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
	//f15 int8
	//f16 int
	//f17 int8
	//f18 int
	//f19 int8
	//f20 int
}

type ItemB struct {
	f1  int8
	f2  int8
	f3  int8
	f4  int8
	f5  int8
	f6  int8
	f7  int8
	f8  int
	f9  int
	f10 int
	f11 int
	f12 int
	f13 int
	f14 int
}

func main() {
	fmt.Println(unsafe.Sizeof(ItemA{}), unsafe.Alignof(ItemA{}))
	fmt.Println(unsafe.Sizeof(ItemB{}), unsafe.Alignof(ItemB{}))
}
