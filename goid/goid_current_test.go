package goid

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

func TestGoid(t *testing.T) {
	var id int64
	for i := 0; i < 100; i++ {
		nid, _ := getGoidByNative()
		sid := getGoidByStack()
		assert.True(t, nid == sid)

		if id == 0 {
			id = sid
		} else {
			assert.True(t, id == sid)
		}
	}
	t.Log(getGoidByStack())
	t.Log(getGoidByNative())
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

	t.Log(getGoidByNative())
}

func TestPC(t *testing.T) {
	const num = 10

	for i := 0; i < num; i++ {
		go func() {
			time.Sleep(time.Minute)
		}()
	}

	time.Sleep(time.Millisecond)
	// can only load pc info, cannot load goroutine context
	res := make([]runtime.StackRecord, 1024)
	n, ok := runtime.GoroutineProfile(res)
	t.Log(n, ok)

	// 获取全部协程上下文快照
	buf := make([]byte, 1<<20)
	n = runtime.Stack(buf, true)
	t.Log("all: \n", string(buf[:n]))
}

// BenchmarkGetGoidByNative-12    	409483110	         2.530 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGetGoidByNative(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getGoidByNative()
	}
}

// BenchmarkGetGoidByStack-12    	  391472	      2996 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGetGoidByStack(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getGoidByStack()
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
