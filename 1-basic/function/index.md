# 函数
```go  
func 函数名(参数)(返回值){
    函数体
}
```
参数由参数变量和参数变量的类型组成，多个参数之间使用英文逗号,分隔，参数中如果相邻变量的类型相同，则可以合并声明只写一个类型。

**1.变长参数**  
Go 语言中可通过在参数名后加...来标识变长参数，变长参数是指函数的参数数量不固定，变长参数通常作为函数的最后一个参数，本质上，函数的变长参数是通过切片来实现的。
```go  
package main
import "fmt"
func Greeting(who ...string) {
	fmt.Printf("%#v", who) //[]string{"Joe", "Anna", "Eileen"}
}
func main() {
	Greeting("Joe", "Anna", "Eileen")
}
```

**2.多个返回值**
返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，Go 语言中函数支持多返回值，多个返回值必须用()包裹，并用英文逗号,分隔；函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过 return 关键字返回。
```go  
func devide(a int, b int) (int, int) {
	var n1, n2 int
	n1 = a / b
	n2 = a % b
	return n1, n2
}

func main() {
	n1, n2 := devide(21, 10)
	fmt.Println(n1, n2)
}
```

# 匿名函数
匿名函数就是没有函数名的函数，当我们不希望给函数起名字的时候，可以使用匿名函数，匿名函数不能够独立存在，它可以被赋值于某个变量或直接对匿名函数进行调用。 

```go  
// 将匿名函数保存到变量
add := func(x, y int) {
    fmt.Printf("The sum of %d and %d is: %d\n", x, y, x+y)
}
add(10, 20) // 通过变量调用匿名函数

// 直接对匿名函数进行调用，最后的一对括号表示对该匿名函数的直接调用执行
func(x, y int) {
    fmt.Printf("The sum of %d and %d is: %d\n", x, y, x+y)
}(10, 20)
```

# 闭包
闭包指的是一个函数和与其相关的引用环境组合而成的实体(即：闭包=函数+引用环境)。
```go  
// 函数 incr 返回了一个函数，返回的这个函数就是一个闭包。这个函数中本身是没有定义变量 i 的，而是引用了它所在的环境（函数incr）中的变量 i。
func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
    // 这里的 i 就成为了一个闭包，闭包对外层词法域变量是引用的，即 i 保存着对 x 的引用。
    i := incr()
    fmt.Println(i()) // 1
    fmt.Println(i()) // 2
    fmt.Println(i()) // 3
    
    // 这里调用了三次 incr()，返回了三个闭包，这三个闭包引用着三个不同的 x，它们的状态是各自独立的
    fmt.Println(incr()()) // 1
    fmt.Println(incr()()) // 1
    fmt.Println(incr()()) // 1
}

```