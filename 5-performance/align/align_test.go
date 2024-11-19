package main

import "testing"

func BenchmarkItemA(b *testing.B) {
	var items [1024]ItemA
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = tmp + int(item.f1)
			tmp = tmp + int(item.f2)
			tmp = tmp + int(item.f3)
			tmp = tmp + int(item.f4)
			tmp = tmp + int(item.f5)
			tmp = tmp + int(item.f6)
			tmp = tmp + int(item.f7)
			tmp = tmp + int(item.f8)
			tmp = tmp + int(item.f9)
			tmp = tmp + int(item.f10)
			tmp = tmp + int(item.f11)
			tmp = tmp + int(item.f12)
			tmp = tmp + int(item.f13)
			tmp = tmp + int(item.f14)
			//tmp = tmp + int(item.f15)
			//tmp = tmp + int(item.f16)
			//tmp = tmp + int(item.f17)
			//tmp = tmp + int(item.f18)
			//tmp = tmp + int(item.f19)
			//tmp = tmp + int(item.f20)
		}
		_ = tmp
	}
}

func BenchmarkItemB(b *testing.B) {
	var items [1024]ItemB
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {

			tmp = tmp + int(item.f1)
			tmp = tmp + int(item.f2)
			tmp = tmp + int(item.f3)
			tmp = tmp + int(item.f4)
			tmp = tmp + int(item.f5)
			tmp = tmp + int(item.f6)
			tmp = tmp + int(item.f7)
			tmp = tmp + int(item.f8)
			tmp = tmp + int(item.f9)
			tmp = tmp + int(item.f10)
			tmp = tmp + int(item.f11)
			tmp = tmp + int(item.f12)
			tmp = tmp + int(item.f13)
			tmp = tmp + int(item.f14)
		}
		_ = tmp
	}
}
