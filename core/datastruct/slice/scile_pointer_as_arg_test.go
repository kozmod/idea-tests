package slice

import (
	"fmt"
	"testing"
)

func TestSlicePointerVsValue(t *testing.T) {
	fmt.Print("1. ")
	slice := []string{"a", "a"}
	func(slice []string) {
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
	fmt.Println()

	fmt.Print("2. ")
	slice = []string{"a", "a"}
	func(slice *[]string) {
		(*slice)[0] = "b"
		(*slice)[1] = "b"
		fmt.Print(*slice)
	}(&slice)
	fmt.Print(slice)
	fmt.Println(" <-- the same result as 1")

	fmt.Print("3. ")
	slice = []string{"a", "a"}
	func(slice []string) {
		slice[0] = "b"
		slice[1] = "b"
		slice = append(slice, "a")
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
	fmt.Println()

	fmt.Print("4. ")
	slice = []string{"a", "a"}
	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
	fmt.Println(" <-- slice change array (diff res than 3)")

	fmt.Print("5. ")
	slice = make([]string, 2, 3)
	func(slice []string) {
		slice = append(slice, "a", "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
	fmt.Println(" <-- slice change array (diff res than 3)")

	fmt.Print("6. ")
	slice = make([]string, 1, 3)
	func(slice []string) {
		slice = slice[1:3]
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(len(slice), cap(slice))
		fmt.Print(slice)
	}(slice)
	fmt.Print(len(slice), cap(slice))
	fmt.Println(slice)

	fmt.Print("7. ")
	slice = make([]string, 1, 3)
	func(slice *[]string) {
		*slice = (*slice)[1:3]
		(*slice)[0] = "b"
		(*slice)[1] = "b"
		fmt.Print(len(*slice), cap(*slice))
		fmt.Print(slice)
	}(&slice)
	fmt.Print(len(slice), cap(slice))
	fmt.Print(slice)
	fmt.Println(" <-- slice change array in both pointer (diff res than 6)")
}
