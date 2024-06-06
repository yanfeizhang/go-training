package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "Thiis is a long long long long long long long long long long long long string"

	// 查看字符串内存地址
	fmt.Println(unsafe.Pointer(&s))

	// 查看字符串的StringHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%+v\n", sh)
	fmt.Println(unsafe.Pointer(sh.Data))
	fmt.Println(sh.Len) // 5

	// 查看指针指向的内容
	arrPtr := (*[5]byte)(unsafe.Pointer(sh.Data))
	fmt.Println(*arrPtr)
}
