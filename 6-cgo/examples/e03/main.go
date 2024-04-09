package main

/*
#include <stdint.h>

struct A {
    int i;
    float f;
};

union B {
    int i;
    float f;
};

enum C {
    ONE,
    TWO,
};
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a.f)

	var b C.union_B
	fmt.Println("b.i:", *(*C.int)(unsafe.Pointer(&b)))
	fmt.Println("b.f:", *(*C.float)(unsafe.Pointer(&b)))

	var c C.enum_C = C.TWO
	fmt.Println(c)
	fmt.Println(C.ONE)
	fmt.Println(C.TWO)
}
