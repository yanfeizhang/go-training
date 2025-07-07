package main

import (
	"bytes"
	"io"
	"sync"
	"testing"
)

// 测试对象，模拟一个较大的结构体
type TestObject struct {
	Data    [1024 * 1024]byte
	Buffer  *bytes.Buffer
	Counter int
}

// 不使用内存池
func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(make([]byte, 0, 1024))
		io.WriteString(buf, "Hello, World!")
		_ = buf.String()
	}
}

func BenchmarkObjectWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := &TestObject{
			Buffer: bytes.NewBuffer(make([]byte, 0, 1024)),
		}
		io.WriteString(obj.Buffer, "Hello, World!")
		_ = obj.Buffer.String()
	}
}

// 使用 sync.Pool
var bufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

var objectBufferPool = sync.Pool{
	New: func() interface{} {
		return &TestObject{
			Buffer: bytes.NewBuffer(make([]byte, 0, 1024)),
		}
	},
}

func BenchmarkWithSyncPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Reset()
		io.WriteString(buf, "Hello, World!")
		_ = buf.String()
		bufferPool.Put(buf)
	}
}

func BenchmarkObjectWithSyncPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := objectBufferPool.Get().(*TestObject)
		obj.Buffer.Reset()
		io.WriteString(obj.Buffer, "Hello, World!")
		_ = obj.Buffer.String()
		objectBufferPool.Put(obj)
	}
}

// 使用方法
// go test -bench=. -benchmem
/*
goos: darwin
goarch: amd64
pkg: gotest/tests/test015
cpu: VirtualApple @ 2.50GHz
BenchmarkWithoutPool-12
 8932203	       139.8 ns/op	    1072 B/op	       2 allocs/op
BenchmarkObjectWithoutPool-12     	   16756	     67440 ns/op	 1057843 B/op	       3 allocs/op
BenchmarkWithSyncPool-12          	67363220	        18.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkObjectWithSyncPool-12    	70156462	        18.53 ns/op	       0 B/op	       0 allocs/op
PASS
*/
