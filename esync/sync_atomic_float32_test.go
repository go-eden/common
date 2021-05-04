package esync

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestAtomicFloat32(t *testing.T) {
	var a AtomicFloat32

	for i := 0; i < 100; i++ {
		f := rand.Float32()
		a.Set(f)
		assert.True(t, a.Get() == f)
		assert.True(t, a.Swap(0) == f, a.Get() == 0)
	}
	multiRun(10, func() {
		for i := 0; i < 1000000; i++ {
			a.Set(123.45)
			_ = a.Swap(1111111.1)
			a.Get()
		}
	})
	t.Log(a.String())
}

// BenchmarkAtomicFloat32_Set-12    	188348179	         6.284 ns/op
func BenchmarkAtomicFloat32_Set(b *testing.B) {
	var a AtomicFloat32
	for i := 0; i < b.N; i++ {
		a.Set(float32(i))
	}
}

// BenchmarkAtomicFloat32_Get-12    	1000000000	         0.5498 ns/op
func BenchmarkAtomicFloat32_Get(b *testing.B) {
	var a AtomicFloat32
	for i := 0; i < b.N; i++ {
		_ = a.Get()
	}
}

// BenchmarkAtomicFloat32_Swap-12    	189897183	         6.165 ns/op
func BenchmarkAtomicFloat32_Swap(b *testing.B) {
	var a AtomicFloat32
	for i := 0; i < b.N; i++ {
		_ = a.Swap(float32(i))
	}
}
