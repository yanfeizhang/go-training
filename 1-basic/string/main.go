package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "Thiis is a long long long long long long long long long long long long string"

	// 查看字符串内存地址
	fmt.Println("\n旧字符串的展示")
	fmt.Println(unsafe.Pointer(&s))

	// 查看字符串的StringHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%+v\n", sh)
	fmt.Println(unsafe.Pointer(sh.Data))
	fmt.Println(sh.Len)

	// 查看指针指向的内容
	arrPtr := (*[5]byte)(unsafe.Pointer(sh.Data))
	fmt.Println(*arrPtr)

	// 字符串浅拷贝
	s1 := ""
	p_s1 := &s1
	*p_s1 = s

	fmt.Println("\n浅拷贝字符串的展示")
	fmt.Println(unsafe.Pointer(&s1))
	sh = (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%+v\n", sh)
	fmt.Println(unsafe.Pointer(sh.Data))
	fmt.Println(sh.Len)

	// 字符串深拷贝, 使用 string
	a := "12345"
	b := make([]byte, len(a))
	copy(b, a)
	c := string(b)
	fmt.Println(c)

	// 将指针转为字符串
	fmt.Println("\n新字符串的展示")

	//1) 先将字符串指针解析成 []byte
	slice1 := unsafe.Slice((*byte)(unsafe.Pointer(sh.Data)), len(s))

	//2)将 []byte 转换成 string
	newStr := string(slice1)
	fmt.Println(newStr)

	//3) 查看字符串内存地址，确认进行了深拷贝
	sh3 := (*reflect.StringHeader)(unsafe.Pointer(&newStr))
	fmt.Printf("%+v\n", sh3)
	fmt.Println(unsafe.Pointer(sh3.Data))
	fmt.Println(sh3.Len)

}
