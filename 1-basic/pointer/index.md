# 指针
**指针的声明、赋值、访问**  
有过C编程经验的同学对指针可是太熟系不过了。Golang中的指针和C中概念差不多，不过在使用上要简单不少。

```go  
//声明
var n int = 20
var pn *int

//空指针判断
if nil == pn {
    fmt.Println("这是一个空指针哦！")
}

//指针赋值
pn = &n
fmt.Println("变量n的地址是:", pn)

//使用指针访问值
fmt.Printf("%d\n", *pn)
```

**指针作为函数的参数**  
如果在函数中需要修改参数的值的话，可能就需要通过指针来进行传递了。

```go  
package main

import "fmt"

func swap1(n1 int, n2 int) {
	nTemp := n2
	n2 = n1
	n1 = nTemp
}

func swap2(n1 *int, n2 *int) {
	var nTemp int
	nTemp = *n2
	*n2 = *n1
	*n1 = nTemp
}

func main() {
	var n1 int = 100
	var n2 int = 200
    
    //不使用指针无法在函数中修改n1,n2的值
	swap1(n1, n2)
	fmt.Println("n1:", n1, "\tn2:", n2)

    //使用指针作为参数，才能成功交换n1,n2的值
	swap2(&n1, &n2)
	fmt.Println("n1:", n1, "\tn2:", n2)
}

```

