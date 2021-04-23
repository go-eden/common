package esync

import (
	"fmt"
	"math"
	"sync/atomic"
)

// AtomicBool support atomic operation
type AtomicBool struct {
	value int32
}

func (t *AtomicBool) Set(v bool) {
	if v {
		atomic.StoreInt32(&t.value, 1)
	} else {
		atomic.StoreInt32(&t.value, 0)
	}
}

func (t *AtomicBool) Get() bool {
	return atomic.LoadInt32(&t.value) != 0
}

func (t *AtomicBool) String() string {
	return fmt.Sprint(t.Get())
}

// AtomicFloat32 support atomic operation
type AtomicFloat32 struct {
	value uint32
}

func (t *AtomicFloat32) Set(v float32) {
	atomic.StoreUint32(&t.value, math.Float32bits(v))
}

func (t *AtomicFloat32) Get() float32 {
	return math.Float32frombits(atomic.LoadUint32(&t.value))
}

func (t *AtomicFloat32) String() string {
	return fmt.Sprint(t.Get())
}

// AtomicFloat64 support atomic operation
type AtomicFloat64 struct {
	value uint64
}

func (t *AtomicFloat64) Set(v float64) {
	atomic.StoreUint64(&t.value, math.Float64bits(v))
}

func (t *AtomicFloat64) Get() float64 {
	return math.Float64frombits(atomic.LoadUint64(&t.value))
}

func (t *AtomicFloat64) String() string {
	return fmt.Sprint(t.Get())
}
