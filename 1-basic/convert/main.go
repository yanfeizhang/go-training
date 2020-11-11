package main

import "fmt"

func main() {
	//1.类型强制转换
	var a float32 = 5.6
	var b int = 10
	//fmt.Println(a * b) //直接将a、b相乘会报错
	fmt.Println(a * float32(b)) //显式将类型进行转换以后才可以相乘
}
