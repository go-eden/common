package esync

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestReMutex1(t *testing.T) {
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

func TestReMutex2(t *testing.T) {
	var m ReMutex

	const num = 10
	const max = 500

	var doneCond = sync.NewCond(&m)
	var conds [num]*sync.Cond
	var counter int32

	for i := 0; i < num; i++ {
		conds[i] = sync.NewCond(&m)
	}

	for i := 0; i < num; i++ {
		var group = int32(i)
		go func() {
			//gid := goid.Gid()
			for atomic.LoadInt32(&counter) < max {
				m.Lock()
				if atomic.LoadInt32(&counter)%num == group {
					//t.Log(gid, counter)
					newCounter := atomic.AddInt32(&counter, 1)
					conds[newCounter%num].Signal() // notify next goroutine
				} else {
					conds[group].Wait()
				}
				m.Unlock()
			}
			doneCond.Signal()
		}()
	}
	m.Lock()
	doneCond.Wait()
	m.Unlock()
}

func TestReMutex3(t *testing.T) {

}

// BenchmarkRemutex-12    	39420012	        30.5 ns/op	       0 B/op	       0 allocs/op
func BenchmarkRemutex(b *testing.B) {
	var m ReMutex

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		m.Lock()
		m.Lock()
		m.Lock()
		m.Unlock()
		m.Lock()
		m.Unlock()
		m.Unlock()
		m.Lock()
		m.Unlock()
		m.Unlock()
	}
}
