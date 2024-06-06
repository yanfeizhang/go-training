package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Device struct {
	Id    int64
	Phone string
	Imei  string
}

type MyStruct struct {
	Id1 int
	Id2 int
}

func main() {
	//创建一个切片
	s := make([]Device, 0, 5)
	s = append(s, Device{
		Id:    10000,
		Phone: "13822222222",
		Imei:  "967029040684350",
	})

	s = append(s, Device{
		Id:    10000,
		Phone: "13822222222",
		Imei:  "967029040684350",
	})

	s = append(s, Device{
		Id:    10000,
		Phone: "18938303322",
		Imei:  "879718719874831738",
	})
	fmt.Println(len(s), cap(s), s)

	// 查看切片内存地址
	fmt.Println(unsafe.Pointer(&s))

	// 查看切片Header
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("%+v\n", sliceHeader)
	fmt.Println(unsafe.Pointer(sliceHeader.Data))
	fmt.Println(sliceHeader.Len)
	fmt.Println(sliceHeader.Cap)

	// 查看指针指向的内容
	devicePtr := (*Device)(unsafe.Pointer(sliceHeader.Data))
	fmt.Println(devicePtr)
	fmt.Println(*devicePtr)

	fmt.Println(unsafe.Pointer(&devicePtr.Phone))
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&devicePtr.Phone))
	fmt.Printf("%+v\n", stringHeader)
	fmt.Println(unsafe.Pointer(stringHeader.Data))
	fmt.Println(stringHeader.Len)

	// 强制将一段内存空间解释为slice
	var a1 = [16]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	a2 := unsafe.Slice((*MyStruct)(unsafe.Pointer(&a1)), 8)
	fmt.Printf("a1: (%v) (%v)\n", unsafe.Pointer(&a1), a1)
	fmt.Printf("a2: (%v) (%v)\n", unsafe.Pointer(&a2), a2)

	sliceHeader = (*reflect.SliceHeader)(unsafe.Pointer(&a2))
	fmt.Printf("%+v\n", sliceHeader)
	fmt.Println(unsafe.Pointer(sliceHeader.Data))
	fmt.Println(sliceHeader.Len)
	fmt.Println(sliceHeader.Cap)

	// 查看字符串内存地址
	/*fmt.Println(unsafe.Pointer(&s))

	// 查看字符串的StringHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%+v\n", sh)
	fmt.Println(unsafe.Pointer(sh.Data))
	fmt.Println(sh.Len) // 5

	// 查看指针指向的内容
	arrPtr := (*[5]byte)(unsafe.Pointer(sh.Data))
	fmt.Println(*arrPtr)*/
}
