package main

import "fmt"

func main() {
	//遍历切片
	s := []int{1, 2, 3, 4, 5}
	for k, v := range s {
		fmt.Println(k, v)
	}

	//遍历map
	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2
	m["key3"] = 3
	for k, v := range m {
		fmt.Println(k, v)
	}

	//switch的使用
	n := 2
	switch n {
	case 1:
		fmt.Println("case 1 hit!")
	case 2:
		fmt.Println("case 2 hit!")
	default:
		fmt.Println("no case hit!")
	}

}
