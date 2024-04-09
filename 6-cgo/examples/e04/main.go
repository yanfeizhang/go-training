package main

/*
#include <string.h>
char arr[10];
char *s = "Hello";
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 通过 reflect.SliceHeader 转换
	var arr0 []byte
	var arr0Hdr = (*reflect.SliceHeader)(unsafe.Pointer(&arr0))
	arr0Hdr.Data = uintptr(unsafe.Pointer(&C.arr[0]))
	arr0Hdr.Len = 10
	arr0Hdr.Cap = 10

	// 通过切片语法转换
	arr1 := (*[31]byte)(unsafe.Pointer(&C.arr[0]))[:10:10]
	fmt.Println(arr1)

	var s0 string
	var s0Hdr = (*reflect.StringHeader)(unsafe.Pointer(&s0))
	s0Hdr.Data = uintptr(unsafe.Pointer(C.s))
	s0Hdr.Len = int(C.strlen(C.s))

	sLen := int(C.strlen(C.s))
	s1 := string((*[31]byte)(unsafe.Pointer(C.s))[:sLen:sLen])
	fmt.Println(s1)
}
