package goid

import (
	"runtime"
)

const (
	ptrSize   = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const
	stackSize = 1024
)

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
