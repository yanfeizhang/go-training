# 类型断言
Golang中有一种interface{}可以代表任意类型，使用起来非常灵活。不过我们有的时候需要知道它到底包含的是啥类型的对象，这时候就需要用到类型断言了

```go  
//类型断言用法1
var c interface{} = 10
switch c.(type) {
case int:
    fmt.Println("int")
case float32:
    fmt.Println("string")
}

//类型断言用法2
var d interface{} = 10
t1, ok := d.(int)
if ok {
    fmt.Println("int", t1)
}
t2, ok := d.(float32)
if ok {
    fmt.Println("float32", t2)
}
```
