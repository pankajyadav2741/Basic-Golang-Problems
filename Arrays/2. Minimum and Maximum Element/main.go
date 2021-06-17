package main

import (
	"fmt"
)

func findMinMax(arr []int) (bool, int, int) {
	var min, max int
	arrLen := len(arr)

	if arrLen == 0 {
		return true, 0, 0
	} else if arrLen == 1 {
		return false, arr[0], arr[0]
	}

	if arr[0] > arr[1] {
		max = arr[0]
		min = arr[1]
	} else {
		min = arr[0]
		max = arr[1]
	}
	for i := 2; i < arrLen; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	return false, min, max
}

func main() {
	arr := []int{10, 2, 45, 67, 125, -10, -27}
	isEmpty, min, max := findMinMax(arr)
	if isEmpty {
		fmt.Printf("Empty Array. Minimum and Maximum element cannot be found.")
	} else {
		fmt.Printf("Minimum Element = %v, Maximum Element = %v\n", min, max)
	}
}
