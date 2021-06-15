package emath

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

func TestMinInt(t *testing.T) {
	assert.True(t, MinInt(math.MaxInt64) == math.MaxInt64)
	assert.True(t, MinInt(math.MaxInt64, 0) == 0)
	assert.True(t, MinInt(1, 2, 3, 10, -1, -100, -200, 300) == -200)

	arr := make([]int, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Int()
	}

	min := MinInt(arr...)

	for _, v := range arr {
		assert.True(t, v >= min)
	}
}

func TestMaxInt(t *testing.T) {
	assert.True(t, MaxInt(math.MinInt64) == math.MinInt64)
	assert.True(t, MaxInt(math.MinInt64, 0) == 0)

	arr := make([]int, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Int()
	}

	max := MaxInt(arr...)

	for _, v := range arr {
		assert.True(t, v <= max)
	}
}

func TestByteMinMax(t *testing.T) {
	arr := make([]byte, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = byte(rand.Int())
	}
	max := MaxByte(arr...)
	min := MinByte(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func TestInt8MinMax(t *testing.T) {
	arr := make([]int8, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = int8(rand.Int())
	}
	max := MaxInt8(arr...)
	min := MinInt8(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func TestUint8MinMax(t *testing.T) {
	arr := make([]uint8, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = uint8(rand.Int())
	}
	max := MaxUint8(arr...)
	min := MinUint8(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func TestInt16MinMax(t *testing.T) {
	arr := make([]int16, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = int16(rand.Int())
	}
	max := MaxInt16(arr...)
	min := MinInt16(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func TestUint16MinMax(t *testing.T) {
	arr := make([]uint16, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = uint16(rand.Int())
	}
	max := MaxUint16(arr...)
	min := MinUint16(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func TestInt32MinMax(t *testing.T) {
	arr := make([]int32, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = int32(rand.Int())
	}
	max := MaxInt32(arr...)
	min := MinInt32(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func TestUint32MinMax(t *testing.T) {
	arr := make([]uint32, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = uint32(rand.Int())
	}
	max := MaxUint32(arr...)
	min := MinUint32(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func TestInt64MinMax(t *testing.T) {
	arr := make([]int64, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = int64(rand.Int())
	}
	max := MaxInt64(arr...)
	min := MinInt64(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func TestUint64MinMax(t *testing.T) {
	arr := make([]uint64, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = uint64(rand.Int())
	}
	max := MaxUint64(arr...)
	min := MinUint64(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func TestFloat32MinMax(t *testing.T) {
	arr := make([]float32, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = float32(rand.Int())
	}
	max := MaxFloat32(arr...)
	min := MinFloat32(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func TestFloat64MinMax(t *testing.T) {
	arr := make([]float64, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = float64(rand.Int())
	}
	max := MaxFloat64(arr...)
	min := MinFloat64(arr...)
	for _, v := range arr {
		assert.True(t, v <= max)
		assert.True(t, v >= min)
	}
}

func BenchmarkMinMax(b *testing.B) {
	var i1 = rand.Float64()
	var i2 = rand.Float64()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MinFloat64(i1, i2)
		MaxFloat64(i1, i2)
	}
}

func BenchmarkMathMinMax(b *testing.B) {
	var i1 = rand.Float64()
	var i2 = rand.Float64()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		math.Min(i1, i2)
		math.Max(i1, i2)
	}
}

// BenchmarkMinMax-12        	218639050	         5.09 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMathMinMax-12    	267942950	         4.54 ns/op	       0 B/op	       0 allocs/op
