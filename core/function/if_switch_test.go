package function

import (
	"fmt"
	"testing"
)

func TestForIf(t *testing.T) {
	for i := 0; i < 7; i++ {
		if i == 2 {
			break
		}
		fmt.Printf("%d\n", i)
	}
	fmt.Println("end")
}

func TestForSwitch(t *testing.T) {
	end := 0
	for x := 0; x <= 7; x++ {
		fmt.Printf("%d - ", x)
		switch {
		case x == 0:
			fmt.Println("start")
		case x == 1:
			fmt.Println("one")
		case x == 2:
			fmt.Println("two")
			break // goto A
		case x > 2:
			fmt.Println("more then two")
		default:
			fmt.Println("counting")
		}
		// A
		end = x
	}
	fmt.Printf("End: %d\n", end)
}

func TestForSwitch2(t *testing.T) {
	end := 0
loop:
	for x := 0; x <= 7; x++ {
		fmt.Printf("%d - ", x)
		switch {
		case x == 0:
			fmt.Println("start")
		case x == 1:
			fmt.Println("one")
		case x == 2:
			fmt.Println("two")
			break loop //goto B
		case x > 2:
			fmt.Println("more then two")
		default:
			fmt.Println("counting")
		}
		end = x
	}
	// B
	fmt.Printf("End: %d\n", end)
}

func TestForSwitch3(t *testing.T) {
	for x := 0; x <= 2; x++ {
		switch {
		case x == 0:
			fmt.Println("start")
		case x == 1:
			fallthrough
		case x == 2:
			fmt.Println("two")
		default:
			fmt.Println("counting")
		}
	}
}
