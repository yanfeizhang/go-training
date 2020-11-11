package main

import "fmt"

func main() {
	//类型断言
	//1类型断言用法1
	var c interface{} = 10
	switch c.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("string")
	}
	//2类型断言用法2
	var d interface{} = 10
	t1, ok := d.(int)
	if ok {
		fmt.Println("int", t1)
	}
	t2, ok := d.(float32)
	if ok {
		fmt.Println("float32", t2)
	}
}
