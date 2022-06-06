package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func TestPprof(t *testing.T) {
	f, _ := os.OpenFile("fib.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()

	n := 10
	for i := 1; i <= 5; i++ {
		fmt.Printf("fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}

	pprof.Lookup("goroutine").WriteTo(f, 0)
}