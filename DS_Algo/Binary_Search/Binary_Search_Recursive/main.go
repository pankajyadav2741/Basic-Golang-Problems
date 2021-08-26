package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	elem := 5
	found := binarySearch(arr, elem, 0, len(arr)-1)
	if found {
		fmt.Printf("Element %v found in array\n", elem)
	} else {
		fmt.Printf("Element %v NOT found in array\n", elem)
	}
}

func binarySearch(arr []int, elem int, left int, right int) bool {
	if left > right {
		return false
	}
	mid := (left + right) / 2
	if elem == arr[mid] {
		return true
	} else if elem < arr[mid] {
		return binarySearch(arr, elem, left, mid-1)
	}
	return binarySearch(arr, elem, mid+1, right)
}
