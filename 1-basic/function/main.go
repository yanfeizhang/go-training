package main

import (
	"fmt"
)

func Greeting(who ...string) {
	fmt.Printf("%#v", who) //[]string{"Joe", "Anna", "Eileen"}
}

func devide(a int, b int) (int, int) {
	var n1, n2 int
	n1 = a / b
	n2 = a % b
	return n1, n2
}

// 函数 incr 返回了一个函数，返回的这个函数就是一个闭包。这个函数中本身是没有定义变量 i 的，而是引用了它所在的环境（函数incr）中的变量 i。
func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	//1.变长参数
	Greeting("Joe", "Anna", "Eileen")

	//2.多返回值函数
	n1, n2 := devide(21, 10)
	fmt.Println(n1, n2)

	//3.匿名函数
	//3.1将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Printf("The sum of %d and %d is: %d\n", x, y, x+y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//3.2直接对匿名函数进行调用，最后的一对括号表示对该匿名函数的直接调用执行
	func(x, y int) {
		fmt.Printf("The sum of %d and %d is: %d\n", x, y, x+y)
	}(10, 20)

	//4. 闭包
	//4.1 这里的 i 就成为了一个闭包，闭包对外层词法域变量是引用的，即 i 保存着对 x 的引用。
	i := incr()
	fmt.Println(i()) // 1
	fmt.Println(i()) // 2
	fmt.Println(i()) // 3

	//4.2 这里调用了三次 incr()，返回了三个闭包，这三个闭包引用着三个不同的 x，它们的状态是各自独立的
	fmt.Println(incr()()) // 1
	fmt.Println(incr()()) // 1
	fmt.Println(incr()()) // 1
}
