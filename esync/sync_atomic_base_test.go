package esync

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

/////////////////////////////////////////////////////////////////////////////////////////
func TestAtomicBool(t *testing.T) {
	var a AtomicBool

	for i := 0; i < 100; i++ {
		b := i%2 == 0
		a.Set(b)
		assert.True(t, a.Get() == b)
		assert.True(t, a.Swap(true) == b && a.Get() == true)
	}
	multiRun(10, func() {
		for i := 0; i < 1000000; i++ {
			a.Set(true)
			_ = a.Swap(false)
			a.Get()
		}
	})
	t.Log(a.String())
}

// BenchmarkAtomicBool_Set-12    	216363957	         5.621 ns/op
func BenchmarkAtomicBool_Set(b *testing.B) {
	var a AtomicBool
	for i := 0; i < b.N; i++ {
		a.Set(i&1 == 1)
	}
}

// BenchmarkAtomicBool_Get-12    	1000000000	         0.2763 ns/op
func BenchmarkAtomicBool_Get(b *testing.B) {
	var a AtomicBool
	for i := 0; i < b.N; i++ {
		_ = a.Get()
	}
}

// BenchmarkAtomicBool_Swap-12    	185458958	         6.179 ns/op
func BenchmarkAtomicBool_Swap(b *testing.B) {
	var a AtomicBool
	for i := 0; i < b.N; i++ {
		_ = a.Swap(i&1 == 1)
	}
}

/////////////////////////////////////////////////////////////////////////////////////////
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

/////////////////////////////////////////////////////////////////////////////////////////
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
