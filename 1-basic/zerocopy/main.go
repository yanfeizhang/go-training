package main

import "fmt"

import (
	"reflect"
	"unsafe"
)

// StringToString 零拷贝复制字符串（实际上 Go 的字符串赋值 already 是零拷贝的）
func StringToString(s string) string {
	// 获取原字符串的底层结构
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	// 创建一个新的字符串头，复制指针和长度
	var newStr string
	newStrHeader := (*reflect.StringHeader)(unsafe.Pointer(&newStr))
	newStrHeader.Data = strHeader.Data + 4 // 往后挪四个字节
	newStrHeader.Len = strHeader.Len - 4

	return newStr
}

// BytesToString 零拷贝将[]byte转换为字符串
func BytesToString(b []byte) string {
	// 获取字节切片的底层结构指针
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	// 构造字符串的底层结构，复用切片的Data和Len
	strHeader := reflect.StringHeader{
		Data: sliceHeader.Data + 4,
		Len:  sliceHeader.Len - 4,
	}

	// 将字符串结构指针转换为string
	return *(*string)(unsafe.Pointer(&strHeader))
}

func main() {
	// 1、[]byte -> string 的零拷贝测试
	var b1 = []byte("hello123234349493053405")
	slice1Header := (*reflect.SliceHeader)(unsafe.Pointer(&b1))
	fmt.Printf("byte array 1：%+v\n", slice1Header)
	fmt.Println(unsafe.Pointer(slice1Header.Data))
	fmt.Println(slice1Header.Len)
	fmt.Println(slice1Header.Cap)

	// 1) 普通拷贝
	l := len(b1)
	s1 := string(b1[4:l])
	string1Header := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Printf("byte array 2：%+v\n", string1Header)
	fmt.Println(unsafe.Pointer(string1Header.Data))
	fmt.Println(string1Header.Len)
	//fmt.Println(string1Header.Cap)
	fmt.Println(s1)

	// 2) 零拷贝
	s2 := BytesToString(b1)
	string2Header := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Printf("byte array 2：%+v\n", string2Header)
	fmt.Println(unsafe.Pointer(string2Header.Data))
	fmt.Println(string2Header.Len)
	fmt.Println(s2)

	// 2. String 的 零拷贝测试
	/*var s2 = "1234567890xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxuyyyyyy"

	// 查看 s2 的header
	string2Header := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Printf("string 1：%+v\n", string2Header)
	fmt.Println(unsafe.Pointer(string2Header.Data))
	fmt.Println(string2Header.Len)

	s3 := StringToString(s1)

	// 查看s2 的header
	string3Header := (*reflect.StringHeader)(unsafe.Pointer(&s3))
	fmt.Printf("string 2：%+v\n", string3Header)
	fmt.Println(unsafe.Pointer(string3Header.Data))
	fmt.Println(string3Header.Len)
	fmt.Println(s3) */
}
