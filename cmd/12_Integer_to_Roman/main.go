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
	// Slice of symbols, including subtractive forms
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// Slice of associated numbers
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	// Empty initial result
	result := ""

	// Offset in the slice of numbers,
	// don't iterate from the beginning of the slice every time
	offset := 0

	// Iterate until the input number is zero
	for num != 0 {
		// Iterate over the shifted slice of numbers
		for i, cur := range numbers[offset:] {
			// Check if the input number is greater than the current number
			if num >= cur {
				// Shift the offset, so we don't have to check for larger numbers next time
				offset += i

				// Iterate until the input number is less than the current number,
				// repeat the same symbol several times if necessary
				for num >= cur {
					// Add the associated symbol to the result
					result += symbols[offset]

					// Subtract the input number
					num -= cur
				}

				// Stop iterating over the slice of numbers after am associated symbol has been found
				break
			}
		}
	}

	return result
}
