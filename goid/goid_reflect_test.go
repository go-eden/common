package goid

import "testing"

func TestGetGoidByReflect(t *testing.T) {
	t.Log(getGoidByReflect())
	t.Log(getGoidByReflect())
	t.Log(getGoidByReflect())
}

// BenchmarkGetGoidByReflect-12    	886767097	         1.20 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGetGoidByReflect(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getGoidByReflect()
	}
}
