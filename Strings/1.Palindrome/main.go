package main

import (
	"fmt"
)

func palindrome(str string) bool {
	strLen := len(str)
	match := true
	for i := 0; i < strLen/2; i++ {
		if str[i] != str[strLen-i-1] {
			match = false
		}
	}
	return match
}

func main() {
	//str := "hello"
	str := "madam"
	isPalin := palindrome(str)
	if isPalin == true {
		fmt.Printf("String %s is palindrome", str)
	} else {
		fmt.Printf("String %s is NOT palindrome", str)
	}
}
