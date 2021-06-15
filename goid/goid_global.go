package goid

import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

const (
	ptrSize   = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const
	stackSize = 1024
)

//go:linkname runtimeG runtime.g
type runtimeG struct {
}

//go:linkname runtimeAtomicAllG runtime.atomicAllG
func runtimeAtomicAllG() (**runtimeG, uintptr)

// getAllGoidByNative retrieve all goid through runtime.atomicAllG
func getAllGoidByNative() (goids []int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("get all goid failed: %v", e))
			goids = nil
		}
	}()
	root, n := runtimeAtomicAllG()
	goids = make([]int64, n)
	for i := uintptr(0); i < n; i++ {
		ptr := *(**runtimeG)(unsafe.Pointer(uintptr(unsafe.Pointer(root)) + i*ptrSize))
		gid := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + offset))
		goids[i] = *gid
	}
	return
}

// getAllGoidByStack find all goid through stack,
// WARNING: This function could be very inefficient.
func getAllGoidByStack() (goids []int64) {
	count := runtime.NumGoroutine()
	size := count * stackSize // it's ok?
	buf := make([]byte, size)
	n := runtime.Stack(buf, true)
	buf = buf[:n]
	// parse all goids
	goids = make([]int64, 0, count+4)
	for i := 0; i < len(buf); {
		goid, off := findNextGoid(buf, i)
		if goid > 0 {
			goids = append(goids, goid)
		}
		i = off
	}
	return
}
