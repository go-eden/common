package goid

import (
	"testing"
	"time"
)

func TestAllGoid(t *testing.T) {
	const num = 10
	for i := 0; i < num; i++ {
		go func() {
			time.Sleep(time.Second)
		}()
	}
	time.Sleep(time.Millisecond)

	nids, err := getAllGoidByNative()
	t.Log("all native goids: ", len(nids), nids, err)

	sids := getAllGoidByStack()
	t.Log("all sids: ", len(sids), sids)
}

// When routineNum = 16
// BenchmarkGetAllGoidByNative-12    	 6226236	       242.9 ns/op	     896 B/op	       1 allocs/op
// When routineNum = 64
// BenchmarkGetAllGoidByNative-12    	 2064740	       740.3 ns/op	    3072 B/op	       1 allocs/op
// When routineNum = 256
// BenchmarkGetAllGoidByNative-12    	  402352	      3163 ns/op	    9472 B/op	       1 allocs/op
// When routineNum = 1024
// BenchmarkGetAllGoidByNative-12    	   84907	     18508 ns/op	   40960 B/op	       1 allocs/op
// When routineNum = 8192
// BenchmarkGetAllGoidByNative-12    	    7520	    253149 ns/op	  270336 B/op	       1 allocs/op
// When routineNum = 65536
// BenchmarkGetAllGoidByNative-12    	     464	   4181609 ns/op	 1581056 B/op	       1 allocs/op
func BenchmarkGetAllGoidByNative(b *testing.B) {
	const routineNum = 65536
	for i := 0; i < routineNum; i++ {
		go func() {
			time.Sleep(time.Minute)
		}()
	}
	time.Sleep(time.Millisecond * 100)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = getAllGoidByNative()
	}
}

// When routineNum = 16
// BenchmarkGetAllGoidByStack-12    	    7779	    147593 ns/op	   57793 B/op	       2 allocs/op
// When routineNum = 64
// BenchmarkGetAllGoidByStack-12    	    2754	    463350 ns/op	  206595 B/op	       2 allocs/op
// When routineNum = 256
// BenchmarkGetAllGoidByStack-12    	     843	   1858839 ns/op	  801156 B/op	       2 allocs/op
// When routineNum = 1024
// BenchmarkGetAllGoidByStack-12    	     254	   6923118 ns/op	 3181195 B/op	       2 allocs/op
// When routineNum = 8192
// BenchmarkGetAllGoidByStack-12    	      43	  37269949 ns/op	16924678 B/op	       2 allocs/op
// When routineNum = 65536
// BenchmarkGetAllGoidByStack-12    	       6	 316648415 ns/op	135282688 B/op	       2 allocs/op
func BenchmarkGetAllGoidByStack(b *testing.B) {
	const routineNum = 65536
	for i := 0; i < routineNum; i++ {
		go func() {
			time.Sleep(time.Minute)
		}()
	}
	time.Sleep(time.Millisecond * 100)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = getAllGoidByStack()
	}
}
