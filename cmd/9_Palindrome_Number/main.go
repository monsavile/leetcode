package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func isPalindrome1(x int) bool {
	s := strconv.Itoa(x)

	sr := strings.Split(s, "")
	slices.Reverse(sr)

	r := strings.Join(sr, "")

	return s == r
}

func isPalindrome2(x int) bool {
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
