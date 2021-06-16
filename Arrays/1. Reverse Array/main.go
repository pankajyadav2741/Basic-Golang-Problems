package main

import (
	"fmt"
)

func main() {
	arr := [5]int {1,2,3,4,5}
	arrLen := len(arr)
	fmt.Println("Before Reversal:", arr)
	for i:=0; i<arrLen/2; i++ {
		arr[i], arr[arrLen-i-1] = arr[arrLen-i-1], arr[i]
	}
	fmt.Println("After Reversal:", arr)
}
