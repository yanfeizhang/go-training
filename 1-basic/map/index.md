# Map

**1.map的定义与元素添加**  
使用make来创建map。  
```go  
m1 := make(map[string]int)
m1["key1"] = 1
m1["key2"] = 2
fmt.Println(m1) //map[key1:1 key2:2]
```

**2.判断key是否存在**  
通过 key 作为索引下标来访问 map，如果 key 不存在，那么将得到 value 对应类型的零值，比如nil、''、false 和 0，取值操作总会有值返回。判断key是否存在可以通过第二个参数来判断，如下：  
```go  
value, ok := m1["key2"]
if ok {
    fmt.Println("key2", value)
}
```
**3.map元素删除**  
golang提供了内建函数 delete() 从 map 中删除一组键值对。  
```go  
delete(m1, "key2")
```
**4.map元素遍历**  
使用 for range 遍历 map，遍历的顺序是随机的，每一次遍历的顺序都不相同，与添加键值对的顺序无关。
```go  
for k, v := range m1 {
    fmt.Println(k, v)
}
```
