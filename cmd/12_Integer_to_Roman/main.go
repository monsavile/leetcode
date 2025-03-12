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
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""
	offset := 0

	for num > 0 {
		for i, v := range value[offset:] {
			if num >= v {
				offset += i
				result += symbol[offset]
				num -= v
				break
			}
		}
	}

	return result
}
