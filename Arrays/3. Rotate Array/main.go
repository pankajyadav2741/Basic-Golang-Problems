package main

import (
	"fmt"
)

func main() {
	arrA := [5]int {1,2,3,4,5}
	arrLen := len(arrA)
	arrB := make([]int, arrLen, arrLen)
	numRotate := 2
	//numRotate := 0
	//numRotate := 5
	if numRotate == 0 || numRotate == arrLen {
		fmt.Println("Array already rotated")
	} else {
		for n:=1; n<=numRotate; n++ {
			fmt.Println("arrA = ", arrA)
			arrB[0]=arrA[arrLen-1]
			for i:=0; i<arrLen-1; i++ {
				fmt.Printf("Num=%d, i=%d, arrB=%v, arrA=%v\n",n, i, arrB, arrA)
				arrB[i+1] = arrA[i]
				
			}
			fmt.Printf("Array after rotation: %v\n", arrB)
			fmt.Println("arrA = ", arrA)
			arrA = arrB
			fmt.Println("-----------")
		}
		fmt.Printf("Array after %d times rotation: %v\n", numRotate, arrB)
	}
}

