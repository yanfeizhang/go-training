package main

import "testing"

func BenchmarkValue(b *testing.B) {
	var a = ItemA{}
	for i := 0; i < b.N; i++ {
		someValueFunc(a)
	}
}

func BenchmarkPointer(b *testing.B) {
	var a = ItemA{}
	for i := 0; i < b.N; i++ {
		somePointerFunc(&a)
	}
}
