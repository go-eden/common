package etime

import (
	"github.com/go-eden/common/esync"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInterval(t *testing.T) {
	var counter esync.AtomicUint16
	var it = NewInterval(time.Millisecond*100, func() {
		counter.Inc()
	})
	time.Sleep(time.Second)

	cv := counter.Get()
	assert.True(t, cv == 10 || cv == 9, cv)
	it.Close()
	assert.True(t, counter.Get() == cv)
}
