
## []byte 到 string 的零拷贝

一般我们使用 string() 将 []byte 转换为字符串
```
var b1 = []byte("hello123234349493053405")
s1 := string(b1[4:l])
```

但这样设计到了内存拷贝，如果数组很大的话会浪费内存申请、内存拷贝开销。
以下是一个零拷贝的函数

```
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
```

拷贝后，通过分别查看 slice 和 string 的header，可以确认其 Data 指向的是同一块区域。

``` 
slice1Header := (*reflect.SliceHeader)(unsafe.Pointer(&b1))
fmt.Printf("byte array 1：%+v\n", slice1Header)
fmt.Println(unsafe.Pointer(slice1Header.Data))
fmt.Println(slice1Header.Len)
fmt.Println(slice1Header.Cap)

string1Header := (*reflect.StringHeader)(unsafe.Pointer(&s1))
fmt.Printf("byte array 2：%+v\n", string1Header)
fmt.Println(unsafe.Pointer(string1Header.Data))
fmt.Println(string1Header.Len)
```