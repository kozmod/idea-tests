package internal

import (
	"fmt"
	"time"
)

func multiplyTwo(v int) int {
	time.Sleep(1 * time.Second)
	return v * 2
}

func square(v int) int {
	time.Sleep(2 * time.Second)
	return v * v
}
func addQuoute(v int) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("'%d'", v)
}
func addFoo(v string) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("%s - Foo", v)
}
