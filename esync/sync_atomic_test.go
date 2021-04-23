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
		a.Set(i%2 == 0)
		assert.True(t, a.Get() == (i%2 == 0))
	}

	signs := make(chan bool, 2)
	go func() {
		for i := 0; i < 1000000; i++ {
			a.Set(true)
			a.Get()
		}
		signs <- true
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			a.Set(false)
			a.Get()
		}
		signs <- true
	}()
	<-signs
	<-signs
	t.Log(a.String())
}

// BenchmarkAtomicBool_Set-12    	244932280	         4.863 ns/op
func BenchmarkAtomicBool_Set(b *testing.B) {
	var a AtomicBool
	for i := 0; i < b.N; i++ {
		a.Set(true)
	}
}

// BenchmarkAtomicBool_Get-12    	1000000000	         0.2763 ns/op
func BenchmarkAtomicBool_Get(b *testing.B) {
	var a AtomicBool
	for i := 0; i < b.N; i++ {
		_ = a.Get()
	}
}

/////////////////////////////////////////////////////////////////////////////////////////
func TestAtomicFloat32(t *testing.T) {
	var a AtomicFloat32

	for i := 0; i < 100; i++ {
		f := rand.Float32()
		a.Set(f)
		assert.True(t, a.Get() == f)
	}

	signs := make(chan bool, 2)
	go func() {
		for i := 0; i < 1000000; i++ {
			a.Set(123.45)
			a.Get()
		}
		signs <- true
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			a.Set(2345.67)
			a.Get()
		}
		signs <- true
	}()
	<-signs
	<-signs
	t.Log(a.String())
}

// BenchmarkAtomicFloat32_Set-12    	226838348	         5.043 ns/op
func BenchmarkAtomicFloat32_Set(b *testing.B) {
	var a AtomicFloat32
	for i := 0; i < b.N; i++ {
		a.Set(123.45)
	}
}

// BenchmarkAtomicFloat32_Get-12    	1000000000	         0.5498 ns/op
func BenchmarkAtomicFloat32_Get(b *testing.B) {
	var a AtomicFloat32
	for i := 0; i < b.N; i++ {
		_ = a.Get()
	}
}

/////////////////////////////////////////////////////////////////////////////////////////
func TestAtomicFloat64(t *testing.T) {
	var a AtomicFloat64

	for i := 0; i < 1000; i++ {
		f := rand.Float64()
		a.Set(f)
		assert.True(t, a.Get() == f)
	}

	signs := make(chan bool, 2)
	go func() {
		for i := 0; i < 1000000; i++ {
			a.Set(rand.Float64())
			a.Get()
		}
		signs <- true
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			a.Set(rand.Float64())
			a.Get()
		}
		signs <- true
	}()
	<-signs
	<-signs
	t.Log(a.String())
}

// BenchmarkAtomicFloat64_Set-12    	241627206	         4.923 ns/op
func BenchmarkAtomicFloat64_Set(b *testing.B) {
	var a AtomicFloat64
	for i := 0; i < b.N; i++ {
		a.Set(123.45)
	}
}

// BenchmarkAtomicFloat64_Get-12    	1000000000	         0.5715 ns/op
func BenchmarkAtomicFloat64_Get(b *testing.B) {
	var a AtomicFloat64
	for i := 0; i < b.N; i++ {
		_ = a.Get()
	}
}
