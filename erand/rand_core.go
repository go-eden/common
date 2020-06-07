package erand

import (
	"github.com/go-eden/common/etime"
	"math/rand"
)

var gr = rand.New(rand.NewSource(etime.NowMicrosecond()))

// Bytes return [n]byte which contains random data.
func Bytes(n int) []byte {
	data := make([]byte, n, n)
	gr.Read(data)
	return data
}
