压测命令

```sh
# go test -cpu=1 -benchtime=2s -benchmem -bench .
```


输出结果

```sh
goos: darwin
goarch: arm64
pkg: go-training/5-performance/string
BenchmarkPlusConcat    	      90	  26262490 ns/op	530995432 B/op	    9999 allocs/op
BenchmarkSprintfConcat 	      51	  46342234 ns/op	832113008 B/op	   34028 allocs/op
BenchmarkBuilderConcat 	   43982	     55024 ns/op	  514800 B/op	      23 allocs/op
BenchmarkBufferConcat  	   58647	     41524 ns/op	  368576 B/op	      13 allocs/op
BenchmarkByteConcat    	   50816	     47209 ns/op	  621296 B/op	      24 allocs/op
BenchmarkPreByteConcat 	  100256	     24414 ns/op	  212992 B/op	       2 allocs/op
PASS
ok  	go-training/5-performance/string	18.174s
```