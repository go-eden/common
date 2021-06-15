package goid

import (
	"fmt"
	"regexp"
)

var pattern = regexp.MustCompile(`goroutine (\d+) \[`)

// Find the next goid from `buf`
func findNextGoid(buf []byte) (goid int64) {
	var matches = pattern.FindSubmatch(buf)
	if l := len(matches); l != 2 {
		if l != 0 {
			fmt.Println("should never be here, any bug happens")
		}
		return
	}
	return str2num(matches[1])
}

// Find all goids from `buf`
func findAllGoid(buf []byte) (goids []int64) {
	arr := pattern.FindAllSubmatch(buf, len(buf))
	goids = make([]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) != 2 {
			fmt.Println("should never be here, any bug happens")
			continue
		}
		goids[i] = str2num(arr[i][1])
	}
	return
}

// Should be faster
func str2num(buf []byte) (v int64) {
	for i := 0; i < len(buf); i++ {
		switch buf[i] {
		case '0':
			v *= 10
		case '1':
			v = v*10 + 1
		case '2':
			v = v*10 + 2
		case '3':
			v = v*10 + 3
		case '4':
			v = v*10 + 4
		case '5':
			v = v*10 + 5
		case '6':
			v = v*10 + 6
		case '7':
			v = v*10 + 7
		case '8':
			v = v*10 + 8
		case '9':
			v = v*10 + 9
		default:
			fmt.Println("should never be here, any bug happens")
		}
	}
	return
}
