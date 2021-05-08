package esync

import (
	"sync/atomic"
)

type AtomicPointer struct {
	p atomic.Value
}

func (t *AtomicPointer) Get() interface{} {
	if v := t.p.Load(); v != nil {
		return *(v.(*interface{}))
	}
	return nil
}

func (t *AtomicPointer) Set(v interface{}) {
	t.p.Store(&v)
}
