package linkname_test

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"
	"time"

	_ "unsafe"
)

// Event types in the trace, args are given in square brackets.
const (
	traceEvGoBlock = 20 // goroutine blocks [timestamp, stack]
)

type mutex struct {
	// Futex-based impl treats it as uint32 key,
	// while sema-based impl as M* waitm.
	// Used to be a union, but unions break precise GC.
	key uintptr
}

//go:linkname lock runtime.lock
func lock(l *mutex)

//go:linkname unlock runtime.unlock
func unlock(l *mutex)

//go:linkname goparkunlock runtime.goparkunlock
func goparkunlock(lock *mutex, reason string, traceEv byte, traceskip int)

func Test(t *testing.T) {
	l := &mutex{}
	go func() {
		lock(l)
		goparkunlock(l, "test", traceEvGoBlock, 1)
	}()
	for i := 0; i < 1; i++ {
		fmt.Printf("---> iteration %d\n", i)
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		time.Sleep(time.Second * 1)
	}
}
