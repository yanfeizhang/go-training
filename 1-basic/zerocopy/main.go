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

// ZeroCopyBytesToString: zero copy []byte->string
func ZeroCopyBytesToString(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	strHeader := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&strHeader))
	//return unsafe.String(unsafe.SliceData(b), len(b))
}

// ZeroCopyBytesToBytes: zero copy []byte->[]byte
func ZeroCopyBytesToBytes(b []byte) []byte {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	newSliceHeader := reflect.SliceHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
		Cap:  sliceHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&newSliceHeader))
}

func ZeroCopyBytesToBytes121(b []byte) []byte {
	return unsafe.Slice(unsafe.SliceData(b), len(b))
}

// 验证这样使用没发生拷贝，估计runtime做过优化
func BytesParam(b []byte) {
	slice1Header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("origin byte array：%+v\n", slice1Header)
	fmt.Printf("\t%p\n", unsafe.Pointer(slice1Header.Data))
	fmt.Printf("\t%d\n", slice1Header.Len)
	fmt.Printf("\t%d\n", slice1Header.Cap)
}

func main() {
	// 1、[]byte -> string 的零拷贝测试
	var b1 = []byte("123456789012345678901234567890")
	slice1Header := (*reflect.SliceHeader)(unsafe.Pointer(&b1))
	fmt.Printf("origin byte array：%+v\n", slice1Header)
	fmt.Printf("\t%p\n", unsafe.Pointer(slice1Header.Data))
	fmt.Printf("\t%d\n", slice1Header.Len)
	fmt.Printf("\t%d\n", slice1Header.Cap)
	fmt.Printf("\n\n")
	fmt.Println(b1)

	// 1) 普通拷贝
	//l := len(b1)
	s1 := string(b1[4:16])
	string1Header := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Printf("non zero copied string ：%+v\n", string1Header)
	fmt.Printf("\t%p\n", unsafe.Pointer(string1Header.Data))
	fmt.Printf("\t%d\n", string1Header.Len)
	fmt.Printf("\t%s\n", s1)

	// 2) 零拷贝
	s2 := ZeroCopyBytesToString(b1[4:16])
	string2Header := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Printf("zero copied：%+v\n", string2Header)
	fmt.Printf("\t%p\n", unsafe.Pointer(string2Header.Data))
	fmt.Printf("\t%d\n", string2Header.Len)
	fmt.Println(s2)
	fmt.Printf("\n\n")

	// 2. []byte->[]byte
	// 1）普通拷贝
	b2 := []byte(string(b1[4:16]))
	slice2Header := (*reflect.SliceHeader)(unsafe.Pointer(&b2))
	fmt.Printf("non zero copied byte array b2：%+v\n", slice2Header)
	fmt.Printf("\t%p\n", unsafe.Pointer(slice2Header.Data))
	fmt.Printf("\t%d\n", slice2Header.Len)
	fmt.Printf("\t%d\n", slice2Header.Cap)
	fmt.Println(b2)

	// 2）zero copy
	b3 := ZeroCopyBytesToBytes(b1[4:16])
	slice3Header := (*reflect.SliceHeader)(unsafe.Pointer(&b3))
	fmt.Printf("zero copied byte array b3 ：%+v\n", slice3Header)
	fmt.Printf("\t%p\n", unsafe.Pointer(slice3Header.Data))
	fmt.Printf("\t%d\n", slice3Header.Len)
	fmt.Printf("\t%d\n", slice3Header.Cap)
	fmt.Println(b3)

	// 3) zero copy: 这个方法不太行，data/len是对的，但是Cap沿用旧的是不行的
	b4 := b1[4:16]
	slice4Header := (*reflect.SliceHeader)(unsafe.Pointer(&b4))
	fmt.Printf("zero copied byte array b3 ：%+v\n", slice4Header)
	fmt.Printf("\t%p\n", unsafe.Pointer(slice4Header.Data))
	fmt.Printf("\t%d\n", slice4Header.Len)
	fmt.Printf("\t%d\n", slice4Header.Cap)
	fmt.Println(b4)

	// 4) zero copy: 这个方法可行
	b5 := ZeroCopyBytesToBytes121(b1[4:16])
	slice5Header := (*reflect.SliceHeader)(unsafe.Pointer(&b5))
	fmt.Printf("zero copied byte array b5 ：%+v\n", slice5Header)
	fmt.Printf("\t%p\n", unsafe.Pointer(slice5Header.Data))
	fmt.Printf("\t%d\n", slice5Header.Len)
	fmt.Printf("\t%d\n", slice5Header.Cap)
	fmt.Println(b5)
	fmt.Printf("\n\n")

	// 3.验证 []byte[offset:] 会不会发生拷贝
	BytesParam(b1[4:16])

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
