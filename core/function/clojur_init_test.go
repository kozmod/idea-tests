package function

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"
	"time"
)

type call struct {
	time time.Time
}

func (call *call) doThx() time.Time {
	return call.time
}

type testStruct struct {
	call func()
}

func (t *testStruct) doThx() {
	t.call()
}

func newTestStruct(time time.Time) *testStruct {
	c := &call{time}
	return &testStruct{
		call: func() {
			fmt.Println(c.time)
		},
	}
}

//go tools trace trace.out
func Test(t *testing.T) {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	now := time.Now()
	ts := newTestStruct(now)
	for i := 0; i < 1_000; i++ {
		ts.doThx()
		time.Sleep(1 * time.Microsecond)
	}
}
