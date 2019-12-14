package etime

import (
	"syscall"
	"time"
)

// NowSecond obtains the current second, use syscall for better performance
func NowSecond() uint32 {
	var tv syscall.Timeval
	if err := syscall.Gettimeofday(&tv); err != nil {
		return uint32(time.Now().Unix())
	}
	return uint32(tv.Sec)
}

// NowMillisecond obtains the current millisecond, use syscall for better performance
func NowMillisecond() int64 {
	var tv syscall.Timeval
	if err := syscall.Gettimeofday(&tv); err != nil {
		return time.Now().UnixNano() / 1e6
	}
	return int64(tv.Sec)*1e3 + int64(tv.Usec)/1e3
}

// NowMicrosecond obtains the current microsecond, use syscall for better performance
func NowMicrosecond() int64 {
	var tv syscall.Timeval
	if err := syscall.Gettimeofday(&tv); err != nil {
		return time.Now().UnixNano() / 1e3
	}
	return int64(tv.Sec)*1e6 + int64(tv.Usec)
}
