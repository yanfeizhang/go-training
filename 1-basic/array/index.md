# 数组
**1.数组的定义**  
var 数组变量名 [元素数量]类型  

```go  
var arr [10]int
```

**2.数组的初始化**  
```go  
var a [5]int                  //数组会初始化所有元素为 int 类型的零值，[0 0 0 0 0]
var b = [5]int{1, 2, 3, 4, 5}
var c = [...]int{1, 2, 3}
```

**3.数组的遍历**
```go  
var a = [5]int{1, 2, 3, 4, 5}
for i := 0; i < len(a); i++ {
	fmt.Println(a[i])
}
var b = [5]int{1, 2, 3, 4, 5}
for index, value := range b {
	fmt.Println(index, value)
}
```

注意： 数组是值类型，赋值和传参会复制整个数组，因此只改变副本的值，不会改变数组本身的值，想要在函数内部修改数组的值，可通过指针来传递数组参数。

# 切片

因为数组的长度是固定的，因此在 Go 语言中很少直接使用数组。  

slice 是一个拥有相同类型元素的可变长度的序列，它是基于数组类型做的一层封装，功能更灵活，支持自动扩容和收缩。  

切片是一个引用类型，底层引用一个数组对象，它的内部结构包含指针址、长度和容量，指针指向第一个 slice 元素对应的底层数组元素的地址（slice 的第一个元素并不一定就是数组的第一个元素），长度对 应slice 中元素的数目，容量一般是从 slice 的开始位置到底层数据的结尾位置，长度不能超过容量。  

可以通过使用内置的 len() 函数求长度，使用内置的 cap() 函数求切片的容量。  

**1.切片的基本定义与初始化**
```go  
var s0 []int
s1 := []int{1, 2, 3}
fmt.Println(len(s0), cap(s0)) //0 0
fmt.Println(len(s1), cap(s1)) //3 3
```

**2.通过make创建切片,并用空值初始化**
```go  
s2 := make([]int, 3, 5)
fmt.Println(len(s2), cap(s2), s2) //3 5 [0 0 0]
```
	
**3.空切片判断**
```go  
var s3 []int
if 0 == len(s3) {
    fmt.Println("This is a empty slice")
}
```

**4.为切片添加元素**
```go  
s4 := make([]int, 0, 5)
s4 = append(s4, 1)
s4 = append(s4, 1)
s4 = append(s4, 1)
fmt.Println(len(s4), cap(s4), s4) //3 5 [1 1 1]
```
**5.切片浅拷贝**  
在C++中有浅拷贝的概念，我们这里沿用过来非常便于理解。

```go  
s5 := []int{1, 2, 3}
s6 := s5
s6[0] = 5
fmt.Println(s5) //[5 2 3]

//浅拷贝情况2
s7 := []int{1, 2, 3}
s8 := s7[1:]
s8[0] = 5
fmt.Println(s7) //[1 5 3]
fmt.Println(s8) //[5 3]
```

**6.切片深拷贝**
```go  
s9 := []int{1, 2, 3}
s10 := make([]int, 2, 5)
copy(s10, s9) //注意：只会拷贝到目标切片的len长度，超过的丢弃
s7[0] = 5
fmt.Println(s9)  //[1 2 3]
fmt.Println(s10) //[1 2]
```

**7.切片删除**  
golang没有直接的删除功能，可以使用append来模拟

```go  
s11 := []int{1, 2, 3, 4, 5}
s11 = append(s11[:2], s11[3:]...)
fmt.Println(s11) //s11:[1 2 4 5]
```
	
