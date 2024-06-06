# 字符串类型
`var a string = 'I am a string!'`


# 字符串内存布局

一个字符串是由两个部分组成，一个是StringHeader，

```
type StringHeader struct {
	Data uintptr
	Len  int
}
```

例如对于以下字符串：

```
func main() {
	s := "Thiis is a long long long long long long long long long long long long string"
}
```

![](string.png)

其StringHeader的值为
- Data：0x10aad09   指向数据段中的字符串
- Len：77


# 参考
- [Go源码学习: string的内部数据结构](https://blog.frognew.com/2021/11/read-go-sources-string.html)