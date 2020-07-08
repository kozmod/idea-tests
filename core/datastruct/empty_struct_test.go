package datastruct_test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/magiconair/properties/assert"
)

func TestNotAllocateMemory(t *testing.T) {
	a := struct{}{}
	size := unsafe.Sizeof(a)
	println(size)
	assert.Equal(t, uintptr(0), size)
}

func TestImplementSet(t *testing.T) {
	set := make(map[string]struct{})
	for _, value := range []string{"apple", "orange", "apple"} {
		set[value] = struct{}{}
	}
	fmt.Println(set)
}

func TestSignalToMain(t *testing.T) {
	worker := func(ch chan struct{}) {
		// Receive a message from the main program.
		<-ch
		println("worker")

		// Send a message to the main program.
		close(ch)
	}

	ch := make(chan struct{})
	go worker(ch)

	// Send a message to a worker.
	ch <- struct{}{}

	// Receive a message from the worker.
	<-ch
	println("end")
}
