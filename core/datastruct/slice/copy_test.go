package slice

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	// Creating slices
	slc1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	var slc2 []int
	slc3 := make([]int, 5)
	slc4 := []int{78, 50, 67, 77}

	// Before copying
	fmt.Println("Slice_1:", slc1)
	fmt.Println("Slice_2:", slc2)
	fmt.Println("Slice_3:", slc3)
	fmt.Println("Slice_4:", slc4)

	// Copying the slices
	copy_1 := copy(slc2, slc1)
	fmt.Println("\nSlice:", slc2)
	fmt.Println("Total number of elements copied:", copy_1)

	copy_2 := copy(slc3, slc1)
	fmt.Println("\nSlice:", slc3)
	fmt.Println("Total number of elements copied:", copy_2)

	copy_3 := copy(slc4, slc1)
	fmt.Println("\nSlice:", slc4)
	fmt.Println("Total number of elements copied:", copy_3)

	// Don't confuse here, because in above
	// line of code the slc4 has been copied
	// and hence modified permanently i.e.
	// slc 4 contains [58 69 40 45]
	copy_4 := copy(slc1, slc4)
	fmt.Println("\nSlice:", slc1)
	fmt.Println("Total number of elements copied:", copy_4)
}

func TestCopy2(t *testing.T) {
	// source slice
	slice_1 := []string{"Geeks", "for", "Geeks", "GFG"}

	// creating destination slice
	// using make function
	slice_2 := make([]string, 3)

	// Before Copying
	fmt.Println("Slice_1: ", slice_1)
	fmt.Println("Slice_2: ", slice_2)

	// Copying slice_1 into slice_2
	// using copy function
	Copy_1 := copy(slice_2, slice_1)
	fmt.Println("\nSlice_1: ", slice_1)
	fmt.Println("Slice_2: ", slice_2)
	fmt.Println("Number of elements copied: ", Copy_1)

	// Copying the slice
	// using copy function
	// see the code clearly
	Copy_2 := copy(slice_1, []string{"123geeks", "gfg"})
	fmt.Println("\nSlice_1 : ", slice_1)
	fmt.Println("Number of elements copied:", Copy_2)
}
