package esync

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAtomicUint16(t *testing.T) {
	var num AtomicUint16

	assert.True(t, num.Inc() == 1)
	assert.True(t, num.Add(2) == 3)
	num.Set(100)
	assert.True(t, num.Swap(200) == 100)
	assert.True(t, num.Get() == 200)

	num.Set(0)
	for i := 1; i < 200000; i++ {
		assert.True(t, num.Add(1) == uint16(i&0xFFFF))
	}
	multiRun(5, func() {
		for i := 0; i < 200000; i++ {
			num.Set(100)
			_ = num.Swap(200)
			num.Get()
		}
	})
}

// BenchmarkAtomicUint16-12    	255484669	         4.639 ns/op
func BenchmarkAtomicUint16(b *testing.B) {
	var num AtomicUint16
	for i := 0; i < b.N; i++ {
		num.Add(i)
	}
}
