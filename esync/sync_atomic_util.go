package esync

import (
	"math"
	"sync/atomic"
	"unsafe"
)

func LoadFloat32(x *float32) float32 {
	return math.Float32frombits(atomic.LoadUint32((*uint32)(unsafe.Pointer(x))))
}

func StoreFloat32(x *float32, v float32) {
	atomic.StoreUint32((*uint32)(unsafe.Pointer(x)), math.Float32bits(v))
}

func SwapFloat32(x *float32, v float32) float32 {
	return math.Float32frombits(atomic.SwapUint32((*uint32)(unsafe.Pointer(x)), math.Float32bits(v)))
}

func LoadFloat64(x *float64) float64 {
	return math.Float64frombits(atomic.LoadUint64((*uint64)(unsafe.Pointer(x))))
}

func StoreFloat64(x *float64, v float64) {
	atomic.StoreUint64((*uint64)(unsafe.Pointer(x)), math.Float64bits(v))
}

func SwapFloat64(x *float64, v float64) float64 {
	return math.Float64frombits(atomic.SwapUint64((*uint64)(unsafe.Pointer(x)), math.Float64bits(v)))
}
