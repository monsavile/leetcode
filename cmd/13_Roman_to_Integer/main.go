package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	symbols := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	result := 0
	offset := 0

	letters := strings.Split(s, "")

	for offset < len(letters) {
		cur := letters[offset]

		if offset < len(letters)-1 {
			next := letters[offset+1]

			if num, ok := symbols[cur+next]; ok {
				result += num
				offset += 2
				continue
			}
		}

		if v, ok := symbols[cur]; ok {
			result += v
		}

		offset++
	}

	return result
}
