package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}

func longestCommonPrefix(strs []string) string {
	// First word
	first := strs[0]

	// Iterate over each letter in the first word
	for idx := range len(first) {
		// Iterate over other words
		for _, current := range strs[1:] {
			// Check if part of the current word matches part of the first word
			if idx == len(current) || current[idx] != first[idx] {
				return first[:idx]
			}
		}
	}

	return first
}
