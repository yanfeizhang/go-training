package main

import "testing"

func benchmark(b *testing.B, f func(int64) string) {
	var roomid int64 = 23423423423423496
	for i := 0; i < b.N; i++ {
		f(roomid)
	}
}

func BenchmarkStrconvConvert(b *testing.B) { benchmark(b, strconvConvert) }
func BenchmarkSprintfConvert(b *testing.B) { benchmark(b, sprintfConvert) }
