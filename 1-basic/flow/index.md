和c语言相比，golang对流程控制语句进行了一些精简，比如省去了c里面的括号。对于if/for都是。

# if/else
基本格式如下:
```go  
if condition {
    分支1
} else if condition {
    分支2
} else {
    分支3
}
```  

# for循环
```go  
for 初始语句;条件表达式;结束语句 {
    循环体语句
}

for i := 0; i < 10; i++ {
	fmt.Println(i)
}
// 无限循环
for {
    ...
}
```

# for/range
在golang中，对于数组、切片、map以及channel等数据结构，用for/range遍历起来会更方便

```go  
//遍历切片
s := []int{1, 2, 3, 4, 5}
for k, v := range s {
    fmt.Println(k, v)
}

//遍历map
m := make(map[string]int)
m["key1"] = 1
m["key2"] = 2
m["key3"] = 3
for k, v := range m {
    fmt.Println(k, v)
}
```

# switch分支
需要特别说一下的是，和c里的switch不同，Go 语言中不需要特别使用 break 语句来表示分支结束。

```go  
n := 2
switch n {
case 1:
    fmt.Println("case 1 hit!")
case 2:
    fmt.Println("case 2 hit!")
default:
    fmt.Println("no case hit!")
}
```
