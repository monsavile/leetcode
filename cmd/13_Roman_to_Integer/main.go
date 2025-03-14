package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}

func romanToInt(s string) int {
	// Map of symbols and associated numbers
	symbols := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	// Initialize the result with the last symbol in the input string
	result := symbols[s[len(s)-1]]

	// Iterate over the range of the input string, excluding the last symbol
	for i := range len(s) - 1 {
		// Check if the current number is greater than the next number
		if symbols[s[i]] < symbols[s[i+1]] {
			// Subtract the current number
			result -= symbols[s[i]]
		} else {
			// Add the current number
			result += symbols[s[i]]
		}
	}

	return result
}
