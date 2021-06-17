package goid

import "testing"

func TestGid(t *testing.T) {
	t.Log(Gid())
}

// BenchmarkGid-12    	267249897	         4.394 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Gid()
	}
}
