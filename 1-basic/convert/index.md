# 类型转换
Golang是强类型语言，而且不会像C、C++那样进行进行隐式类型转换，如下的代码会报错

```go  
package main

import "fmt"

func main() {
    var a float32 = 5.6
    var b int = 10
    fmt.Println (a * b)
}
```

报错如下
```go  
invalid operation: a * b (mismatched types float32 and int)
```

必须进行显式的类型转换才可以相乘
```go  
fmt.Println(a * float32(b))
```

