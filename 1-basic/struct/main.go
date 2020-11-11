package main

import (
	"encoding/json"
	"fmt"
)

type Device struct {
	Phone string `json:"phone" xml:"phone"`
	Imei  string `json:"imei" xml:"imei"`
}

type User struct {
	Id     int    `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Device Device `json:"device" xml:"device"`
}

func main() {
	//1.实例化
	//1.1 实例化方式1
	var s1 Device
	fmt.Printf("%#v\n", s1) //main.Device{Phone:"", Imei:""}

	//1.2 实例化方式2
	var s2 = new(Device)
	fmt.Printf("%#v\n", s2) //&main.Device{Phone:"", Imei:""}

	//2.赋值
	//2.1 普通赋值方法
	var s3 Device
	s3.Phone = "13811111111"
	s3.Imei = "867029040684350"
	fmt.Printf("%#v\n", s3) //{13811111111 867029040684350}

	//2.2 实例化同时赋值1
	s4 := Device{
		Phone: "13822222222",
		Imei:  "967029040684350",
	}
	fmt.Printf("%#v\n", s4) //main.Device{Phone:"13811111111", Imei:"867029040684350"}

	//2.3 实例化同时赋值2
	s5 := Device{
		"13833333333",
		"967029040684350",
	}
	fmt.Printf("%#v\n", s5) //main.Device{Phone:"13833333333", Imei:"967029040684350"}

	//4.结构体嵌套
	//4.1 嵌套结构体赋值方式1
	var s6 User
	s6.Id = 1
	s6.Name = "张三"
	s6.Device.Phone = "13833333333"
	s6.Device.Imei = "967029040684350"
	fmt.Printf("%#v\n", s6) //main.User{Id:1, Name:"张三", Device:main.Device{Phone:"13833333333", Imei:"967029040684350"}}

	//4.2 嵌套结构体赋值方式2
	s7 := User{
		Id:   2,
		Name: "李四",
		Device: Device{
			Phone: "13833333333",
			Imei:  "967029040684350",
		},
	}
	fmt.Printf("%#v\n", s7) //main.User{Id:2, Name:"李四", Device:main.Device{Phone:"13833333333", Imei:"967029040684350"}}

	//5.结构体的标签
	//5.1 json序列化
	data, err := json.Marshal(s7)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	fmt.Printf("%s\n", data) //{"id":2,"name":"李四","device":{"phone":"13833333333","imei":"967029040684350"}}

	//5.2 json反序列化
	var s8 User
	err = json.Unmarshal(data, &s8)
	if err != nil {
		fmt.Println("json unmarshal failed:", err)
		return
	}
	fmt.Printf("%#v\n", s8) //main.User{Id:2, Name:"李四", Device:main.Device{Phone:"13833333333", Imei:"967029040684350"}}
}
