// +build go1.16

package goid

import (
	"errors"
	"fmt"
	"unsafe"
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
