package main

import "testing"

type Item struct {
	id  int
	val [4096]byte
}

func BenchmarkForStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangeStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}

func generateItems(n int) []*Item {
	items := make([]*Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, &Item{id: i})
	}
	return items
}

func BenchmarkForPointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangePointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
