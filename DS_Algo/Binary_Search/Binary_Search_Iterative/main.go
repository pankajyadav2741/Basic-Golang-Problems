package main

import "fmt"

var mid, left, right int

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	elem := 30
	found := binarySearch(arr, elem)
	if found {
		fmt.Printf("Element %v found in array\n", elem)
	} else {
		fmt.Printf("Element %v NOT found in array\n", elem)
	}
}

func binarySearch(arr []int, elem int) bool {
	right = len(arr) - 1
	for left <= right {
		mid = (left + right) / 2
		if elem == arr[mid] {
			return true
		} else if elem < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
