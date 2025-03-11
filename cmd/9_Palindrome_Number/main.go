package main

import (
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := make([]int, 0)
	num := x

	for {
		mod := num % 10
		num = num / 10

		s = append(s, mod)

		if num == 0 {
			break
		}
	}

	x2 := 0

	for i, v := range s {
		x2 += v * int(math.Pow10(len(s)-i-1))
	}

	return x == x2
}
