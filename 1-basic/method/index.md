
Go 语言的方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量，因此方法是一种特殊类型的函数。接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，但是接收者不能是一个接口类型，因为接口是一个抽象定义，但方法是具体实现。

方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。

# 方法的定义
```go  
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是而不是 self、this 之类的命名，例如，Context 类型的接收者变量应该命名为 c。

# 值接收者
在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
原因是当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。

```go  
type User struct {
	Name  string
	Email string
}

//值作为接收者，可以访问对象成员
func (u User) printName1() {
	fmt.Println(u.Name)
}

//值作为接收者， 无法修改对象成员
func (u User) changeName1(name string) {
	u.Name = name
}

func main() {
	var u1 User
	u1.Name = "张三"
	u1.Email = "demo@163.com"

	u1.printName1() //张三

	u1.changeName1("李四")
	fmt.Printf("%#v\n", u1) //main.User{Name:"张三", Email:"demo@163.com"}
}

```
# 指针接收者
在指针类型接收者的方法中也可以获取接收者的成员值，而且可以修改接收者变量本身。

```go  
type User struct {
	Name  string
	Email string
}

//指针作为接收者，也可以访问对象成员
func (u *User) printName2() {
	fmt.Println(u.Name)
}

//指针作为接收者，可以修改对象成员
func (u *User) changeName2(name string) {
	u.Name = name
}

func main() {

	var u1 User
	u1.Name = "张三"
	u1.Email = "demo@163.com"

	u1.printName2() //张三

	u1.changeName2("李四")
	fmt.Printf("%#v\n", u1) //main.User{Name:"李四", Email:"demo@163.com"}
}
```

鉴于性能的原因，接收者是拷贝代价比较大的大对象时，接收者通常会是一个指向接收者类型的指针，比如接收者的类型是结构体时。
另外通常为了保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
  



