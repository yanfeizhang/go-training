package main

import "fmt"

type User struct {
	Name  string
	Email string
}

//值作为接收者，可以访问对象成员
func (u User) printName1() {
	fmt.Println(u.Name)
}

//指针作为接收者，也可以访问对象成员
func (u *User) printName2() {
	fmt.Println(u.Name)
}

//值作为接收者， 无法修改对象成员
func (u User) changeName1(name string) {
	u.Name = name
}

//指针作为接收者，可以修改对象成员
func (u *User) changeName2(name string) {
	u.Name = name
}

func main() {

	var u1 User
	u1.Name = "张三"
	u1.Email = "demo@163.com"

	//值接收者和指针接收者的方法都可以访问对象成员
	u1.printName1() //张三
	u1.printName2() //张三

	//值为接收者的方法无法修改对象成员、指针接收者可以
	u1.changeName1("李四")
	fmt.Printf("%#v\n", u1) //main.User{Name:"张三", Email:"demo@163.com"}
	u1.changeName2("李四")
	fmt.Printf("%#v\n", u1) //main.User{Name:"李四", Email:"demo@163.com"}
}
