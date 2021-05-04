package esync

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestAtomicFloat64(t *testing.T) {
	var a AtomicFloat64

	for i := 0; i < 1000; i++ {
		f := rand.Float64()
		a.Set(f)
		assert.True(t, a.Get() == f)
		assert.True(t, a.Swap(0) == f, a.Get() == 0)
	}
	multiRun(10, func() {
		for i := 0; i < 1000000; i++ {
			a.Set(rand.Float64())
			a.Get()
		}
	})
	t.Log(a.String())
}

// BenchmarkAtomicFloat64_Set-12    	183620686	         6.208 ns/op
func BenchmarkAtomicFloat64_Set(b *testing.B) {
	var a AtomicFloat64
	for i := 0; i < b.N; i++ {
		a.Set(float64(i))
	}
}

// BenchmarkAtomicFloat64_Get-12    	1000000000	         0.5715 ns/op
func BenchmarkAtomicFloat64_Get(b *testing.B) {
	var a AtomicFloat64
	for i := 0; i < b.N; i++ {
		_ = a.Get()
	}
}

// BenchmarkAtomicFloat64_Swap-12    	187972417	         6.227 ns/op
func BenchmarkAtomicFloat64_Swap(b *testing.B) {
	var a AtomicFloat64
	for i := 0; i < b.N; i++ {
		_ = a.Swap(float64(i))
	}
}
