package emem

import (
	"unsafe"
)

//go:linkname memclr runtime.memclrNoHeapPointers
func memclr(ptr unsafe.Pointer, n uintptr)

//go:linkname roundupsize runtime.roundupsize
func roundupsize(size uintptr) uintptr

// Clear clear the specified memory, which is [ptr, ptr+n)
func Clear(ptr unsafe.Pointer, n uintptr) {
	memclr(ptr, n)
}
