package main

import "fmt"

func closeAll() {
	fmt.Println("close all 1 !")
}

func main() {
	fmt.Println("start")

	//defer语句
	//for i := 0; i < 5; i++ {
	//	defer fmt.Println(i)
	//}

	//panic语句，终止正常语句执行，但会执行panic
	//defer func() { fmt.Println("close all") }()
	//panic("ERROR")

	//recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	panic("ERROR")

	fmt.Println("end")
}
