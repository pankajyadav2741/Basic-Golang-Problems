package main

import (
	"fmt"
)

func fibonacci(n int) []int {
	if n <= 0 {
		return nil
	}
	prev, cur := 0, 1
	arr := make([]int, n)
	arr[0], arr[1] = prev, cur
	for i := 2; i < n; i++ {
		next := prev + cur
		arr[i] = next
		prev = cur
		cur = next
	}
	return arr
}

func main() {
	n := 15
	arr := fibonacci(n)
	fmt.Printf("First %d numbers in fibonacci series: %v", n, arr)
}
