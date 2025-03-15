package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(3749)) // MMMDCCXLIX
	fmt.Println(intToRoman(58))   // LVIII
	fmt.Println(intToRoman(1994)) // MCMXCIV
}

func intToRoman(num int) string {
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""
	offset := 0

	for num != 0 {
		for i, current := range numbers[offset:] {
			if num >= current {
				offset += i

				for num >= current {
					result += symbols[offset]
					num -= current
				}

				break
			}
		}
	}

	return result
}
