压测命令
```
#go test -cpu=1 -benchtime=2s -benchmem -bench .
```

输出结果
```
goos: darwin
goarch: arm64
pkg: go-training/5-performance/slice
BenchmarkSliceAppend1 	    1369	   1744203 ns/op	41678080 B/op	      38 allocs/op
BenchmarkSliceAppend2 	    5877	    401816 ns/op	 8003584 B/op	       1 allocs/op
PASS
ok  	go-training/5-performance/slice	8.832s
```

