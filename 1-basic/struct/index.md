
# 结构体

**1.结构体定义**
```go  
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
```

其中：  
- 类型名：表示结构体的名称，在同一个包内不能重复。
- 字段名：表示结构体成员的字段名，结构体中的字段名必须唯一。
- 字段类型：表示结构体字段的具体类型，如果相邻的成员类型如果相同的话可以被合并到一行。

如果结构体成员名字是以大写字母开头的，那么该成员就是可导出的，小写表示仅在定义当前结构体的包中可访问。

**2.结构体实例化**  

结构体必须实例化后才能使用结构体的字段，只有实例化后才会分配内存。实例化方法有两种:  
```go  
var 结构体实例 结构体类型
var 结构体实例 = new(结构体类型)
```

刚实例化完没有赋值的结构体，其成员变量都是对应其类型的零值。  
```go 
type Device struct {
	Phone string
	Imei  string
}
 
//1.实例化方式1
var s1 Device
fmt.Printf("%#v\n", s1) //main.Device{Phone:"", Imei:""}  

//2.实例化方式2
var s2 = new(Device)  
fmt.Printf("%#v\n", s2) //&main.Device{Phone:"", Imei:""}  
```

**3.结构体赋值**  
当结构体实例化完了以后，就可以使用其中的字段并进行赋值了。  
```go  
var s3 Device
s3.Phone = "13811111111"
s3.Imei = "867029040684350"
fmt.Println(s3) //{13811111111 867029040684350} 
```

如果嫌实例化、赋值分开略繁琐的话，可以将这两个步骤合并到一起。如下：  
```go  
//实例化同时赋值1
s4 := Device{
    Phone: "13822222222",
    Imei:  "967029040684350",
}
fmt.Printf("%#v\n", s4) //main.Device{Phone:"13811111111", Imei:"867029040684350"}

//实例化同时赋值2
s5 := Device{
    "13833333333",
    "967029040684350",
}
fmt.Println(s5) //main.Device{Phone:"13822222222", Imei:"967029040684350"}
```

**4.嵌套结构体**
结构体中的成员有可能是另外一个结构体，对于这种嵌套型的结构体来说，它的赋值方法稍微有一点点的不同，但也同样很简单。

```go  
type Device struct {
	Phone string
	Imei  string
}

type User struct {
	Id     int
	Name   string
	Device Device
}

//嵌套结构体赋值方式1
var s6 User
s6.Id = 1
s6.Name = "张三"
s6.Device.Phone = "13833333333"
s6.Device.Imei = "967029040684350"
fmt.Printf("%#v\n", s6) //main.User{Id:1, Name:"张三", Device:main.Device{Phone:"13833333333", Imei:"967029040684350"}}

//嵌套结构体赋值方式2
s7 := User{
    Id:   2,
    Name: "李四",
    Device: Device{
        Phone: "13833333333",
        Imei:  "967029040684350",
    },
}
fmt.Printf("%#v\n", s7) //main.User{Id:2, Name:"李四", Device:main.Device{Phone:"13833333333", Imei:"967029040684350"}}
```  

**5.结构体的标签**  
结构体可以在字段后边定义 Tag，由一对反引号包裹起来。标签最经常是在json处理的地方会用到。
```go  
//json序列化
data, err := json.Marshal(s7)
if err != nil {
    fmt.Println("json marshal failed:", err)
    return
}
fmt.Printf("%s\n", data) //{"id":2,"name":"李四","device":{"phone":"13833333333","imei":"967029040684350"}}

//json反序列化
var s8 User
err = json.Unmarshal(data, &s8)
if err != nil {
    fmt.Println("json unmarshal failed:", err)
    return
}
fmt.Printf("%#v\n", s8) //main.User{Id:2, Name:"李四", Device:main.Device{Phone:"13833333333", Imei:"967029040684350"}}
```


 


