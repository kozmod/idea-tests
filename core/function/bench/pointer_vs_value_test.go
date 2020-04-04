package bench

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime/trace"
	"strings"
	"testing"
)

type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func TestBenchmarkMemoryStack(t *testing.T) {
	cmd := exec.Command(
		"go", "test", "./...",
		"-bench=BenchmarkMemoryStack", "-benchmem", "-run=^$", "-count=10",
	)

	cmd.Stdin = strings.NewReader("bench")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\n %s", out.String())
	if err := ioutil.WriteFile("out/stack.txt", out.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}

	outbanch, err := exec.Command("benchstat", "out/stack.txt").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("benchstat outpur:\n%s", outbanch)
}

// go test ./... -bench=BenchmarkMemoryStack -benchmem -run=^$ -count=10 2>&1 | tee out/stack.txt && benchstat out/stack.txt 2>&1 | tee out/stack-bstat.txt
func BenchmarkMemoryStack(b *testing.B) {
	var s S

	f, err := os.Create("out/stack.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		s = byCopy()
	}

	trace.Stop()

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}

// go test ./... -bench=BenchmarkMemoryHeap -benchmem -run=^$ -count=10 2>&1 | tee out/head.txt && benchstat out/head.txt 2>&1 | tee out/head-bstat.txt
func BenchmarkMemoryHeap(b *testing.B) {
	var s *S

	f, err := os.Create("out/heap.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		s = byPointer()
	}

	trace.Stop()

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}
