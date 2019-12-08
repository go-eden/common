package cgo

import "testing"

func TestNow(t *testing.T) {
	t.Log(Now())
}

// BenchmarkNow-12    	20000000	        82.3 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNow(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Now()
	}
}
