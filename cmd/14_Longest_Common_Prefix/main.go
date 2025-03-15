package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}

func longestCommonPrefix(strs []string) string {
	first := strs[0]

	for idx := range len(first) {
		for _, current := range strs[1:] {
			if idx == len(current) || current[idx] != first[idx] {
				return first[:idx]
			}
		}
	}

	return first
}
