package esync

import (
	"fmt"
	"sync/atomic"
)

// AtomicUint16 support atomic operation
type AtomicUint16 struct {
	value uint32
}

func (t *AtomicUint16) Add(v uint16) {
	atomic.AddUint32(&t.value, uint32(v))
}

func (t *AtomicUint16) Set(v uint16) {
	atomic.StoreUint32(&t.value, uint32(v))
}

func (t *AtomicUint16) Get() uint16 {
	return uint16(atomic.LoadUint32(&t.value))
}

func (t *AtomicUint16) Swap(v uint16) uint16 {
	return uint16(atomic.SwapUint32(&t.value, uint32(v)))
}

func (t *AtomicUint16) String() string {
	return fmt.Sprint(t.Get())
}
