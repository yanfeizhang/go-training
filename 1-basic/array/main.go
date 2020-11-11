package main

import "fmt"

func main() {
	//1. 数组的定义、赋值与遍历
	//1.1 指定数组长度定义
	var a1 = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a1); i++ {
		fmt.Println(i, a1[i])
	}

	//1.2 让编译器推导数组长度
	var a2 = [...]int{1, 2, 3, 4, 5}
	for index, value := range a2 {
		fmt.Println(index, value)
	}

	//2.切片的定义、赋值
	//2.1 切片的基本定义与初始化
	var s0 []int
	s1 := []int{1, 2, 3}
	fmt.Println(len(s0), cap(s0)) //0 0
	fmt.Println(len(s1), cap(s1)) //3 3

	//2.2 通过make创建切片,并用空值初始化（推荐）
	s2 := make([]int, 3, 5)
	fmt.Println(len(s2), cap(s2), s2) //3 5 [0 0 0]

	//2.3 空切片判断
	var s3 []int
	if 0 == len(s3) {
		fmt.Println("This is a empty slice")
	}
	//2.4 为切片添加元素
	s4 := make([]int, 0, 5)
	s4 = append(s4, 1)
	s4 = append(s4, 1)
	s4 = append(s4, 1)
	fmt.Println(len(s4), cap(s4), s4) //3 5 [1 1 1]

	//2.5 切片浅拷贝
	//浅拷贝情况1
	s5 := []int{1, 2, 3}
	s6 := s5
	s6[0] = 5
	fmt.Println(s5) //[5 2 3]

	//浅拷贝情况2
	s7 := []int{1, 2, 3}
	s8 := s7[1:]
	s8[0] = 5
	fmt.Println(s7) //[1 5 3]
	fmt.Println(s8) //[5 3]

	//2.6 切片深拷贝
	s9 := []int{1, 2, 3}
	s10 := make([]int, 2, 5)
	copy(s10, s9) //注意：只会拷贝到目标切片的len长度，超过的丢弃
	s7[0] = 5
	fmt.Println(s9)  //[1 2 3]
	fmt.Println(s10) //[1 2]

	//2.6 切片删除
	s11 := []int{1, 2, 3, 4, 5}
	s11 = append(s11[:2], s11[3:]...)
	fmt.Println(s11) //s11:[1 2 4 5]
}
