package esync

import "testing"

func TestReMutex(t *testing.T) {
	var m ReMutex

	m.Lock()
	m.Lock()

	m.Lock()
	m.Unlock()
	m.Lock()
	m.Unlock()

	m.Unlock()
	m.Unlock()

	for i := 0; i < 1000; i++ {
		m.Lock()
		m.Lock()
		m.Unlock()
		m.Unlock()
	}
}

// BenchmarkRemutex-12    	39420012	        30.5 ns/op	       0 B/op	       0 allocs/op
func BenchmarkRemutex(b *testing.B) {
	var m ReMutex

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		m.Lock()
		m.Lock()
		m.Unlock()
		m.Unlock()
	}
}
