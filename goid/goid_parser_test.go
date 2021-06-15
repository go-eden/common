package goid

import (
	"runtime"
	"runtime/debug"
	"testing"
)

func TestGoidParser(t *testing.T) {
	currStack := debug.Stack()
	t.Log("stack: \n", string(currStack))
	id := findNextGoid(currStack)
	t.Log("## ", id)

	buf := make([]byte, 1<<20)
	runtime.Stack(buf, true)
}

func TestGoidParserAll(t *testing.T) {
	buf := make([]byte, 1<<20)
	n := runtime.Stack(buf, true)
	buf = buf[:n]

	t.Log("stack: \n", string(buf))

	ids := findAllGoid(buf)
	t.Log("All goid: ", len(ids), ids)
}
