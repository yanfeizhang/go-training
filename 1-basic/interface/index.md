Go 语言不是一种 “传统” 的面向对象编程语言：它里面没有类和继承的概念，但是 Go 语言里有非常灵活的 接口 概念，通过它可以实现很多面向对象的特性。

接口是定义了一组需要被实现的方法的抽象类型，但是这些方法不包含（实现）代码， 接口类型是对其它类型行为的抽象和概括，对于需要满足接口的数据类型，它需要实现接口所需的所有方法，一个接口包含两部分：一组接口方法和一个接口类型。

使用接口的最大优势，就是在定义函数时使用接口作为参数，那么在调用函数时，任何实现该接口类型的变量都可以作为函数的参数被传递。


# 1.接口的定义
每个接口由数个方法组成，接口的定义格式如下：
```go  
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```
- 接口名：使用 type 将接口定义为自定义的类型名， Go 语言的接口在命名时，（按照约定，只包含一个方法的）接口的名字由方法名加 [e]r 后缀组成，例如 Printer、Reader、Writer、Logger、Converter 等等。还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，此时接口名以 able 结尾。
- 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包之外的包代码访问。
- 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略：

```go  
type Caller interface {
	call()
}
```

# 2.接口的实现
一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。

```go  
type Iphone6s struct {
}

func (i Iphone6s) call() {
	fmt.Println("使用IPhone 6s打个电话")
}
```

# 3.接口的使用
```go  
//可以使用接口作为参数
func CallToMom(c Caller) {
	c.call()
}
//任意实现了该接口中声明方法的对象都可以作为参数传入
var i Iphone6s
CallToMom(i)
```

# 4.特殊接口之空接口
空接口是指没有定义任何方法的接口。
- 1.因此任何类型都实现了空接口，可以给一个空接口类型的变量赋任何类型的值。
- 2.如果一个函数的参数是空接口类型，那么它也可以接收任何参数

```go  
type Any interface{}

func main() {
    var val Any
    val = 5
    fmt.Printf("type:%T value:%v\n", val)   //type:int value:5
    
    str := "string"
    val = str
    fmt.Printf("type:%T value:%v\n", val)  //type:string value:string
}
```

# 5.特殊接口之接口嵌套
一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
```go  
type ReadWrite interface {
    Read(b Buffer) bool
    Write(b Buffer) bool
}

type Lock interface {
    Lock()
    Unlock()
}

// 接口 File 包含了 ReadWrite 和 Lock 的所有方法，它还额外有一个 Close() 方法
type File interface {
    ReadWrite
    Lock
    Close()
}
```

# 6.测试一个值是否实现了某个接口
这是类型断言中的一个特例：假定 v 是一个值，然后我们想测试它是否实现了 Stringer 接口，可以这样做：

```go  
var i Iphone6s
CallToMom(i)

var e interface{} = i
v, ok := e.(Caller)
if ok {
    fmt.Println("该对象实现了Caller接口", v) 
} else {
    fmt.Println("该对象没有实现Caller接口", v)
}
```

