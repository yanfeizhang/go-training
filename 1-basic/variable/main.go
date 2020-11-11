package main

import "fmt"

//全局变量
//多个变量定义建议用括号括起来
var (
	v1 int
	v2 string
)

func main() {
	v1 = 10
	v2 = "I am v2"

	var v3 string = "I am v3" //定义变量同时初始化
	v4 := "I am v4"           //类型推导，根据字面量判断出v4的类型

	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)
	fmt.Println("v3:", v3)
	fmt.Println("v4:", v4)
}
