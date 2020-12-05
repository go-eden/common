package emem

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestMemclr(t *testing.T) {
	var data [12345]byte

	for i := 0; i < len(data); i++ {
		data[i] = byte(i)
	}
	memclr(unsafe.Pointer(&data), uintptr(len(data)))

	assert.True(t, data[0] == 0)
	assert.True(t, data[10] == 0)
	assert.True(t, data[100] == 0)
	assert.True(t, data[1000] == 0)
	assert.True(t, data[len(data)-1] == 0)
}

func TestRound(t *testing.T) {
	t.Log(roundupsize(1))
	t.Log(roundupsize(12))
	t.Log(roundupsize(123))
	t.Log(roundupsize(1234))
	t.Log(roundupsize(12345))
	t.Log(roundupsize(123456))
}

//
// BenchmarkMemTool
// BenchmarkMemTool-12     	32692742	        36.6 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemTool2
// BenchmarkMemTool2-12    	 1000000	      1039 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemTool3
// BenchmarkMemTool3-12    	 1000000	      1032 ns/op	       0 B/op	       0 allocs/op
//

// BenchmarkMemTool-12    	26022583	        38.7 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMemTool(b *testing.B) {
	var blk [4096]byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Clear(unsafe.Pointer(&blk), 4096)
	}
}

// BenchmarkMemTool2-12    	 1104585	      1045 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMemTool2(b *testing.B) {
	var blk [4096]byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		clearBlockForEach(&blk)
	}
}

// BenchmarkMemTool2-12    	 1104585	      1045 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMemTool3(b *testing.B) {
	var blk [4096]byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for x := 0; x < len(blk); x++ {
			blk[x] = 0
		}
	}
}

func clearBlockForEach(blk *[4096]byte) {
	for x := 0; x < len(blk); x++ {
		blk[x] = 0
	}
}
