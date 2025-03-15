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
	symbols := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := symbols[s[len(s)-1]]

	for i := range len(s) - 1 {
		if symbols[s[i]] < symbols[s[i+1]] {
			result -= symbols[s[i]]
		} else {
			result += symbols[s[i]]
		}
	}

	return result
}
