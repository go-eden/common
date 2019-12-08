package cgo

/*
#include <time.h>
#include <sys/time.h>
#include <stdlib.h>

long now() {
	struct timeval time;
	gettimeofday(&time, NULL);
  	return time.tv_sec * 1000000  +  time.tv_usec;
}
*/
import "C"

func Now() int64 {
	n := C.now()
	return int64(n)
}
