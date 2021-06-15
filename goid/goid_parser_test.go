package goid

import (
	"runtime/debug"
	"testing"
)

func TestGoidParser(t *testing.T) {
	currStack := debug.Stack()
	t.Log("stack: \n", string(currStack))
	id, _ := findNextGoid(currStack, 0)
	t.Log("## ", id)
}
