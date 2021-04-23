package esync

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat32(t *testing.T) {
	var v float32
	StoreFloat32(&v, 100)
	assert.True(t, LoadFloat32(&v) == 100)
	StoreFloat32(&v, 12.00)
	assert.True(t, LoadFloat32(&v) == 12.00)
	swaped := SwapFloat32(&v, 120.0)
	assert.True(t, LoadFloat32(&v) == 120.0 && swaped == 12.00)

	signs := make(chan bool, 2)
	go func() {
		for i := 0; i < 1000000; i++ {
			StoreFloat32(&v, 1000)
		}
		signs <- true
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			_ = LoadFloat32(&v)
		}
		signs <- true
	}()
	<-signs
	<-signs
}

func TestFloat64(t *testing.T) {
	var v float64
	StoreFloat64(&v, 100)
	assert.True(t, LoadFloat64(&v) == 100)
	StoreFloat64(&v, 12.00)
	assert.True(t, LoadFloat64(&v) == 12.00)
	swaped := SwapFloat64(&v, 120.0)
	assert.True(t, LoadFloat64(&v) == 120.0 && swaped == 12.00)

	signs := make(chan bool, 2)
	go func() {
		for i := 0; i < 1000000; i++ {
			StoreFloat64(&v, 1000)
		}
		signs <- true
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			_ = LoadFloat64(&v)
		}
		signs <- true
	}()
	<-signs
	<-signs
}

// BenchmarkStoreFloat32-12    	257130040	         4.490 ns/op
func BenchmarkStoreFloat32(b *testing.B) {
	var v float32
	for i := 0; i < b.N; i++ {
		StoreFloat32(&v, 100.1)
	}
}

// BenchmarkLoadFloat32-12    	1000000000	         0.5146 ns/op
func BenchmarkLoadFloat32(b *testing.B) {
	var v float32
	v = 122.3
	for i := 0; i < b.N; i++ {
		_ = LoadFloat32(&v)
	}
}

// BenchmarkSwapFloat32-12    	245113434	         4.706 ns/op
func BenchmarkSwapFloat32(b *testing.B) {
	var v float32
	for i := 0; i < b.N; i++ {
		_ = SwapFloat32(&v, 100.1)
	}
}

// BenchmarkStoreFloat64-12    	235631870	         4.559 ns/op
func BenchmarkStoreFloat64(b *testing.B) {
	var v float64
	for i := 0; i < b.N; i++ {
		StoreFloat64(&v, 100.1)
	}
}

// BenchmarkLoadFloat64-12    	1000000000	         0.2698 ns/op
func BenchmarkLoadFloat64(b *testing.B) {
	var v float64
	v = 122.3
	for i := 0; i < b.N; i++ {
		_ = LoadFloat64(&v)
	}
}

// BenchmarkSwapFloat64-12    	250339864	         4.538 ns/op
func BenchmarkSwapFloat64(b *testing.B) {
	var v float64
	for i := 0; i < b.N; i++ {
		_ = SwapFloat64(&v, 100.1)
	}
}
