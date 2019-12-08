package main

import (
	"github.com/go-eden/common/etime"
	"time"
)

func main() {
	println(etime.NowSecond())
	println(etime.NowMillisecond())
	println(etime.NowMicrosecond())

	println(time.Now().Unix())
	println(time.Now().UnixNano() / 1e6)
	println(time.Now().UnixNano() / 1e3)
}
