package main

import (
	"fmt"
)

func rotate(arr []int, n int) []int {
	arrLen := len(arr)
	if n <= 0 || n == arrLen || arrLen == 0 {
		// Array already rotated or not required
		return arr
	}
	arr = append(arr[arrLen-n:], arr[:arrLen-n]...)
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	numRotate := 3
	fmt.Printf("Array Before rotation: %v\n", arr)
	arr = rotate(arr, numRotate)
	fmt.Printf("Array after %d rotations: %v\n", numRotate, arr)
}
