package esync

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestAtomicInt(t *testing.T) {
	var num AtomicInt

	assert.True(t, num.Inc() == 1)
	assert.True(t, num.Add(2) == 3)
	num.Set(100)
	assert.True(t, num.Swap(200) == 100)
	assert.True(t, num.Get() == 200)

	num.Set(0)
	for i := 1; i < 200000; i++ {
		assert.True(t, num.Add(1) == i)
	}
	multiRun(5, func() {
		for i := 0; i < 200000; i++ {
			num.Set(100)
			_ = num.Swap(200)
			num.Get()
		}
	})
}

// BenchmarkAtomicInt_Add-12    	251870193	         4.817 ns/op
func BenchmarkAtomicInt_Add(b *testing.B) {
	var num AtomicInt
	for i := 0; i < b.N; i++ {
		num.Add(i)
	}
}

// BenchmarkAtomicInt_Get-12    	1000000000	         0.3167 ns/op
func BenchmarkAtomicInt_Get(b *testing.B) {
	var num AtomicInt
	num.Set(rand.Int())
	for i := 0; i < b.N; i++ {
		num.Get()
	}
}

// BenchmarkAtomicIntAll-12    	128318323	         9.233 ns/op
func BenchmarkAtomicIntAll(b *testing.B) {
	var num AtomicInt

	for i := 0; i < b.N; i++ {
		num.Inc()
		_ = num.Swap(i)
	}
}
