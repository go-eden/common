package etime

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	if NowSecond() != uint32(time.Now().Unix()) {
		t.FailNow()
	}

	ms1 := NowMillisecond()
	ms2 := time.Now().UnixNano() / 1e6
	if ms2 > ms1+1 || ms1 > ms2+1 {
		t.FailNow()
	}

	micro1 := NowMicrosecond()
	micro2 := time.Now().UnixNano() / 1e3
	if micro1 > micro2+1 || micro2 > micro1+1 {
		t.FailNow()
	}
}

func BenchmarkNowSecond(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NowSecond()
	}
}

func BenchmarkNowMillisecond(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NowMillisecond()
	}
}

func BenchmarkNowMicrosecond(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NowMicrosecond()
	}
}

func BenchmarkTimeNow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = time.Now().Unix()
	}
}

// BenchmarkNowSecond-12         	29296284	        39.8 ns/op	       0 B/op	       0 allocs/op
// BenchmarkNowMillisecond-12    	29361312	        40.7 ns/op	       0 B/op	       0 allocs/op
// BenchmarkNowMicrosecond-12    	29742286	        40.1 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTimeNow-12           	15010953	        70.0 ns/op	       0 B/op	       0 allocs/op
