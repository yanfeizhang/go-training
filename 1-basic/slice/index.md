
# slice 的内存布局

```
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

例如对于下面的实例代码

```
s := make([]Device, 0, 5)
s = append(s, Device{
    Id:    10000,
    Phone: "13822222222",
    Imei:  "967029040684350",
})

s = append(s, Device{
    Id:    10000,
    Phone: "13822222222",
    Imei:  "967029040684350",
})

s = append(s, Device{
    Id:    10000,
    Phone: "18938303322",
    Imei:  "879718719874831738",
})
```

其内存布局如下图所示

![](slice.png)

## slice的强制转换
可以使用 unsafe.Slice 将一段连续的内存空间强制转换为 slice。

这个操作原理是给这一段内存空间申请一个 SliceHeader 结构体，然后将这段内存空间的首地址赋值给 SliceHeader.Data 字段，
这样就实现了将一段连续的内存空间强制转换为 slice 的操作。

下面是一个例子

```
var a1 = [16]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
a2 := unsafe.Slice((*MyStruct)(unsafe.Pointer(&a1)), 8)
fmt.Printf("a1: (%v) (%v)\n", unsafe.Pointer(&a1), a1)
fmt.Printf("a2: (%v) (%v)\n", unsafe.Pointer(&a2), a2)

sliceHeader = (*reflect.SliceHeader)(unsafe.Pointer(&a2))
fmt.Printf("%+v\n", sliceHeader)
fmt.Println(unsafe.Pointer(sliceHeader.Data))
```

输出结果如下

```shell
a1: (0x1400013e080) ([1 2 3 4 5 1 2 3 4 5 1 2 3 4 5 0])
a2: (0x1400011e048) ([{1 2} {3 4} {5 1} {2 3} {4 5} {1 2} {3 4} {5 0}])
&{Data:1374390837376 Len:8 Cap:8}
0x1400013e080
```

可见a1数组的地址是0x1400013e080，使用 unsafe.Slice 创建出来的 sliceHeader.Data 的地址也是0x1400011e048。
这个操作在实际开发中非常有用，可以避免大块数组的内存拷贝。
