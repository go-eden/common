package goid

import (
	"github.com/go-eden/common/goid/g"
	"runtime"
	"sync"
	"unsafe"
)

var littleBuf = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 64)
		return &buf
	},
}

// getGoidByStack parse the current goroutine's id from caller stack.
// This function could be very slow(like 3000us/op), but it's very safe.
func getGoidByStack() (goid int64) {
	bp := littleBuf.Get().(*[]byte)
	defer littleBuf.Put(bp)

	b := *bp
	b = b[:runtime.Stack(b, false)]
	goid, _ = findNextGoid(b, 0)
	return
}

// getGoidByNative parse the current goroutine's id from G.
// This function could be very fast(like 1ns/op), but it could be failed.
func getGoidByNative() (int64, bool) {
	if offset == 0 {
		return 0, false
	}
	tmp := g.G()
	if tmp == nil {
		return 0, false
	}
	p := (*int64)(unsafe.Pointer(uintptr(tmp) + offset))
	return *p, true
}
