package main

import (
	"fmt"
	"unsafe"
)

// 子结构体
type Child struct {
	ChildField string
}

// 父结构体，包含子结构体
type Parent struct {
	Child
	ParentField int
}

func main() {
	// 创建一个父结构体的实例
	parent := Parent{
		Child: Child{
			ChildField: "This is a child field",
		},
		ParentField: 10,
	}
	// 获取父结构体的指针
	ptr := &parent

	// 将父结构体指针强制转换为子结构体指针
	childPtr := (*Child)(unsafe.Pointer(ptr))
	// 访问子结构体的字段
	fmt.Println("Child field:", childPtr.ChildField)

	parentPtr := (*Parent)(unsafe.Pointer(childPtr))
	fmt.Println("parent field:", parentPtr.ParentField)
}
