package main

import (
	"testing"
)

func BenchmarkSliceAppend1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SliceAppend1()
	}
}

func BenchmarkSliceAppend2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SliceAppend2()
	}
}
