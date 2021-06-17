package goid

import "github.com/go-eden/routine"

// Gid retrieve the current goroutine's unique id.
// It will try get gid by reflect for better performance,
// and could parse gid from stacktrace for failover supporting.
func Gid() (id int64) {
	return routine.Goid()
}
