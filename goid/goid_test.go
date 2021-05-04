package goid

import (
	"runtime"
	"testing"
)

func TestGetGoidByReflect(t *testing.T) {
	t.Log(getGoidByReflect())
	t.Log(getGoidByReflect())
	t.Log(getGoidByReflect())
}

func TestGoid(t *testing.T) {
	t.Log(getGoidByReflect())
	t.Log(getGoidByStack())
}

func TestGid(t *testing.T) {
	var (
		buf [128]byte
		n   = runtime.Stack(buf[:], false)
	)
	t.Log(string(buf[:n]))

	t.Log(getGoidByStack())
	t.Log(getGoidByStack())
	t.Log(getGoidByStack())
	t.Log(getGoidByStack())
	t.Log(getGoidByStack())

	t.Log(getGoidByReflect())
}

// BenchmarkGetGoidByReflect-12    	450743221	         2.543 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGetGoidByReflect(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getGoidByReflect()
	}
}

// BenchmarkGid-12    	267249897	         4.394 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Gid()
	}
}

// BenchmarkGetGoid-12    	  388618	      3012 ns/op	     144 B/op	       3 allocs/op
func BenchmarkGetGoid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getGoidByStack()
	}
}

// 32: 	BenchmarkStack-12    	  464960	      2487 ns/op	      32 B/op	       1 allocs/op
// 64: 	BenchmarkStack-12    	  476478	      2452 ns/op	      64 B/op	       1 allocs/op
// 128: BenchmarkStack-12    	  474406	      2465 ns/op	     128 B/op	       1 allocs/op
func BenchmarkStack(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf [128]byte
		runtime.Stack(buf[:], false)
	}
}
