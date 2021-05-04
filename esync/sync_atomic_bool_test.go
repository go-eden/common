package esync

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAtomicBool(t *testing.T) {
	var a AtomicBool

	for i := 0; i < 100; i++ {
		b := i%2 == 0
		a.Set(b)
		assert.True(t, a.Get() == b)
		assert.True(t, a.Swap(true) == b && a.Get() == true)
	}
	multiRun(10, func() {
		for i := 0; i < 1000000; i++ {
			a.Set(true)
			_ = a.Swap(false)
			a.Get()
		}
	})
	t.Log(a.String())
}

// BenchmarkAtomicBool_Set-12    	216363957	         5.621 ns/op
func BenchmarkAtomicBool_Set(b *testing.B) {
	var a AtomicBool
	for i := 0; i < b.N; i++ {
		a.Set(i&1 == 1)
	}
}

// BenchmarkAtomicBool_Get-12    	1000000000	         0.2763 ns/op
func BenchmarkAtomicBool_Get(b *testing.B) {
	var a AtomicBool
	for i := 0; i < b.N; i++ {
		_ = a.Get()
	}
}

// BenchmarkAtomicBool_Swap-12    	185458958	         6.179 ns/op
func BenchmarkAtomicBool_Swap(b *testing.B) {
	var a AtomicBool
	for i := 0; i < b.N; i++ {
		_ = a.Swap(i&1 == 1)
	}
}

// ------------------------------------------------------------------------------
func multiRun(concurrency int, f func()) {
	signs := make(chan bool, concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer func() {
				if e := recover(); e != nil {
					fmt.Printf("RUN panic: %v", e)
				}
				signs <- true
			}()
			f()
		}()
	}
	for i := 0; i < concurrency; i++ {
		<-signs
	}
}
