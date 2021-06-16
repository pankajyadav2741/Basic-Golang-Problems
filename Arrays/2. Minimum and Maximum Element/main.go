package main

import (
	"fmt"
)

func main() {
	var min, max int
	//arr := []int{}
	//arr := []int{100}
	arr := [7]int {10,2,45,67,125,-10,-27}
	arrLen := len(arr)
	if arrLen == 0 {
		fmt.Println("Array is empty")
	} else if arrLen == 1 {
		fmt.Println("Minimum and Maximum array element:", arr[0])
	} else {
		if arr[0] > arr[1] {
			max = arr[0]
			min = arr[1]	
		} else {
			min = arr[0]
			max = arr[1]		
		}
		for i:=2; i<arrLen; i++ {
			if arr[i] < min {
				min = arr[i]
			}
			if arr[i] > max {
				max = arr[i]
			}
		}	
		fmt.Printf("Minimum = %v, Maximum = %v\n", min, max)
	}
}
