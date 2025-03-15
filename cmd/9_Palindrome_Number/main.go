package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false
	fmt.Println(isPalindrome(10))   // false
}

func isPalindrome(x int) bool {
	original := x
	reversed := 0

	for x > 0 {
		reversed = 10*reversed + x%10
		x /= 10
	}

	return original == reversed
}
