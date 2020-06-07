package erand

import (
	"math/rand"
	"time"
)

var gr *rand.Rand

func init() {
	seed1 := time.Now().UnixNano()
	seed2 := time.Now().UnixNano()
	seed3 := time.Now().UnixNano()
	src := rand.NewSource(seed1 | seed2 | seed3)
	gr = rand.New(src)
}
