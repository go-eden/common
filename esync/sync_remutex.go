package esync

import (
	"fmt"
	"github.com/go-eden/common/goid"
	"sync"
)

// Reentrant Mutex
type ReMutex struct {
	sync.Mutex
	holder  int64
	counter int32
}

func (t *ReMutex) Lock() {
	var gid = goid.Gid()

	// if held by others, wait...
	if t.holder != gid {
		t.Mutex.Lock()
		t.holder = gid
		t.counter = 1
		return
	}

	// do lock unlocked lock
	if t.counter == 0 {
		t.Mutex.Lock()
		t.counter = 1
		return
	}

	// relock
	t.counter++
}

func (t *ReMutex) Unlock() {
	var gid = goid.Gid()

	if t.holder != gid {
		panic(fmt.Sprintf("Lock conflict, held by g[%v]", t.holder))
	}
	if t.counter > 0 {
		t.counter--
	}
	if t.counter == 0 {
		t.Mutex.Unlock()
	}
}
