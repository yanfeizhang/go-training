CGO 的 C 虚拟包提供了一组函数、用于 Go 语言和 C 语言之间数组和字符串的双向转换。

## 拷贝转换
拷贝转换的原理是通过克隆的方式在 Go 语言和 C 语言的环境之中拷贝数据。会涉及到内存分配和拷贝的开销。

```cgo
C.CString(string) *C.char
```
输入的 Go 字符串，克隆一个 C 语言格式的字符串；返回的字符串由 C 语言的 malloc 函数分配，不使用时需要通过 C 语言的 free 函数释放。

```cgo
func C.CBytes([]byte) unsafe.Pointer
```
C.CBytes 函数的功能和 C.CString 类似，用于从输入的 Go 语言字节切片克隆一个 C 语言版本的字节数组，同样返回的数组需要在合适的时候释放。

```cgo
func C.GoString(*C.char) string
```
C.GoString 用于将从 NULL 结尾的 C 语言字符串克隆一个 Go 语言字符串。

```cgo
func C.GoBytes(unsafe.Pointer, C.int) []byte
```
C.GoBytes 用于从 C 语言数组，克隆一个 Go 语言字节切片。

## Go访问C语言，reflect包转换
reflect 包提供了类型信息，可以在 Go 语言中直接访问 C 语言的内存空间
这个访问过程相当于指针类型强制转换，没有内存拷贝的开销。
参见 main.go

## C访问Go语言
在 C 语言中可以通过 GoString 和 GoSlice 来访问 Go 语言的字符串和切片

