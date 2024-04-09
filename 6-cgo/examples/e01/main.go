package main

/*
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}

struct A {
    int i;
    float f;
};

*/
import "C"
import "fmt"

func main() {
	v := 42
	C.printint(C.int(v))

	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a.f)
}
