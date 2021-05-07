package efmt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	tFormatStr = "hello: name=%s, no=%d, score=%.2f"
	tFormatArg = []interface{}{"James", 10000, 123.456}
)

func TestSliceRef(t *testing.T) {
	b := new(buffer)
	b.writeString("hello")
	b.writeString("world")
	t.Log(string(*b))
}

func TestPrinter(t *testing.T) {
	var p Printer

	v1 := fmt.Sprintf(tFormatStr, tFormatArg...)
	v2 := p.Sprintf(tFormatStr, tFormatArg...)

	assert.True(t, v1 == string(v2))
}

// BenchmarkEfmtSprintf-12    	 3754678	       311.5 ns/op	       0 B/op	       0 allocs/op
func BenchmarkEfmtSprintf(b *testing.B) {
	var p Printer
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = p.Sprintf(tFormatStr, tFormatArg...)
	}
}

// BenchmarkStdSprintf-12     	 3251124	       359.8 ns/op	      48 B/op	       1 allocs/op
func BenchmarkStdSprintf(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf(tFormatStr, tFormatArg...)
	}
}
