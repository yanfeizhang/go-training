package main

import (
	"fmt"
	"unsafe"
)

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

func some_func() {
	fmt.Println("some_func()")
}

func main() {
	//1. 指针的初步使用
	//1.1 指针的声明
	var n int = 20
	var pn *int

	//1.2 空指针判断
	if nil == pn {
		fmt.Println("这是一个空指针哦！")
	}

	//1.3 指针赋值
	pn = &n
	fmt.Println("变量n的地址是:", pn)

	//1.4 使用指针访问值
	fmt.Printf("%d\n", *pn)

	//2. 指针作为函数的参数
	var n1 int = 100
	var n2 int = 200

	//2.1 不使用指针无法在函数中修改n1,n2的值
	swap1(n1, n2)
	fmt.Println("n1:", n1, "\tn2:", n2)

	//2.2 使用指针作为参数，才能成功交换n1,n2的值
	swap2(&n1, &n2)
	fmt.Println("n1:", n1, "\tn2:", n2)

	//3.函数指针
	var pSomeFunc func()
	pSomeFunc = some_func
	fmt.Println(unsafe.Pointer(pSomeFunc))
	//fmt.Println((int64)(pSomeFunc))
}
