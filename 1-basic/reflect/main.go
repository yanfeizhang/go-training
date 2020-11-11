package main

import (
	"fmt"
	"reflect"
)

type UserId int

func main() {
	//1.TypeOf返回的是reflect.Type类型
	//1.1 查看TypeOf返回
	var a int = 64
	t1 := reflect.TypeOf(a)            //注意 TypeOf的参数类型是 interface{}
	fmt.Printf("type:%v\n", t1.Name()) //type:int

	//1.2 区分Type返回值的Name和Kind
	var id UserId = 10
	t2 := reflect.TypeOf(id)
	fmt.Printf("type:%v kind:%v\n", t2.Name(), t2.Kind()) //type:UserId kind:int

	//2.ValueOf返回的是reflect.Value类型
	//2.1 查看ValueOf返回
	var n1 int = 100
	v := reflect.ValueOf(n1)
	fmt.Println(v.Int())

	//2.2 修改ValueOf返回
	var n2 int = 100
	v2 := reflect.ValueOf(&n2)
	fmt.Println(v2) //得到的是一个地址：0xc0000b4048
	v2.Elem().SetInt(200)
	fmt.Println(n2) //输出200，修改成功！
}
