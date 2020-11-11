package main

import "fmt"

//1.接口定义
type Caller interface {
	call()
}

//2.接口实现1
type Iphone6s struct {
}

func (i Iphone6s) call() {
	fmt.Println("使用IPhone 6s打个电话")
}

//3.接口实现2
type HuaWeiP30 struct {
}

func (h HuaWeiP30) call() {
	fmt.Println("使用华为P30去打给电话")
}

//4.接口使用：可以使用接口作为参数
func CallToMom(c Caller) {
	c.call()
}

func main() {
	//5.接口使用：任意实现了该接口中声明方法的对象都可以作为参数传入
	var i Iphone6s
	var h HuaWeiP30
	CallToMom(i)
	CallToMom(h)

	//6.测试某个变量是否实现了某接口，其实也就是相当于判断该变量是否是该接口类型
	var e interface{} = h
	v, ok := e.(Caller)
	if ok {
		fmt.Println("该对象实现了Caller接口", v)
	} else {
		fmt.Println("该对象没有实现Caller接口", v)
	}
}
