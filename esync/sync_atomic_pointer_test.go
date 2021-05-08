package esync

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAtomicPointer(t *testing.T) {
	var p AtomicPointer

	assert.True(t, p.Get() == nil)

	s1 := []byte("hello")
	p.Set(s1)
	t.Log(p.Get())
}

// BenchmarkAtomicPointer_Set-12    	18869185	        55.95 ns/op	      40 B/op	       2 allocs/op
func BenchmarkAtomicPointer_Set(b *testing.B) {
	var p AtomicPointer
	s1 := []byte("hello")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		p.Set(s1)
	}
}

// BenchmarkAtomicPointer_Get-12    	797969247	         1.398 ns/op	       0 B/op	       0 allocs/op
func BenchmarkAtomicPointer_Get(b *testing.B) {
	var p AtomicPointer
	p.Set([]byte("hello"))
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = p.Get().([]byte)
	}
}
