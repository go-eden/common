package goid

import (
	"fmt"
)

var anchor = []byte("goroutine ")

// Find the next goid from `buf[off:]`
func findNextGoid(buf []byte, off int) (goid int64, next int) {
	i := off
	hit := false
	// skip to anchor
	for sb := len(buf) - len(anchor); i < sb; {
		if buf[i] == anchor[0] && buf[i+1] == anchor[1] &&
			buf[i+2] == anchor[2] && buf[i+3] == anchor[3] &&
			buf[i+4] == anchor[4] && buf[i+5] == anchor[5] &&
			buf[i+6] == anchor[6] && buf[i+7] == anchor[7] &&
			buf[i+8] == anchor[8] && buf[i+9] == anchor[9] {
			hit = true
			i += len(anchor)
			break
		}
		for ; i < len(buf) && buf[i] != '\n'; i++ {
		}
		i++
	}
	// return if not hit
	if !hit {
		return 0, len(buf)
	}
	// extract goid
	var done bool
	for ; i < len(buf) && !done; i++ {
		switch buf[i] {
		case '0':
			goid *= 10
		case '1':
			goid = goid*10 + 1
		case '2':
			goid = goid*10 + 2
		case '3':
			goid = goid*10 + 3
		case '4':
			goid = goid*10 + 4
		case '5':
			goid = goid*10 + 5
		case '6':
			goid = goid*10 + 6
		case '7':
			goid = goid*10 + 7
		case '8':
			goid = goid*10 + 8
		case '9':
			goid = goid*10 + 9
		case ' ':
			done = true
			break
		default:
			goid = 0
			fmt.Println("should never be here, any bug happens")
		}
	}
	next = i
	return
}
