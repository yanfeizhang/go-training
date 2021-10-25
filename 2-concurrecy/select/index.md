
select 可同时监听多个channel的读/写，执行 select 时，若只有一个 case 通过，则执行该 case。

```go
select {
    case i := <-ch1:
        fmt.Println("Worker1 job done", i)
    case j := <-ch2:
    ...
}
```

通过通道返回的第二个参数可以感知到通道的关闭

```go
select {
    ...
    case _, ok := <-ch3:
        if ok {
            fmt.Println("Finish all job")
            return
        }
    }
```