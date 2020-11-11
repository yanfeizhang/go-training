package main

import "fmt"

func main() {
	//1 map的定义与元素添加
	m1 := make(map[string]int)
	m1["key1"] = 1
	m1["key2"] = 2
	fmt.Println(m1) //map[key1:1 key2:2]

	//2 判断key是否存在
	value, ok := m1["key2"]
	if ok {
		fmt.Println("key2", value)
	}

	//3 map元素删除
	delete(m1, "key2")
	fmt.Println(m1) //map[key1:1]

	//4 map元素遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
