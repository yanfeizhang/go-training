
# go benchmark用法
go bench 的各个参数
- -cpu：指定使用的 GOMAXPROCS，例如 2
- -benchtime：指定为时间长度（ 5s ），或者是次数（ 50x ）
- -count：指定 benchmark 的轮数，例如 3
- -benchmem：指定查看访存大小、申请内存次数

各个压测实例如下
```sh
go test -bench .
go test -cpu=2,4 -bench .
go test -cpu=1 -benchtime=5s -bench .
go test -cpu=1 -benchtime=2s -benchmem -bench .
```