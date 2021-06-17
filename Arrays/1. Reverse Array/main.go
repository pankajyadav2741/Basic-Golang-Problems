package main

import (
	"fmt"
)

func reverse(arr []int) []int {
	arrLen := len(arr)
	for i:=0; i<arrLen/2; i++ {
		arr[i], arr[arrLen-i-1] = arr[arrLen-i-1], arr[i]
	}
	return arr
}

func main() {
	arr := []int {1,2,3,4,5}
	fmt.Println("Before Reversal: ", arr)
	arr = reverse(arr)
	fmt.Println("After Reversal: ", arr)
}
