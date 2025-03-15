package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // []int{0, 1}
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // []int{1, 2}
	fmt.Println(twoSum([]int{3, 3}, 6))         // []int{0, 1}
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i1, num := range nums {
		if i2, ok := hash[target-num]; ok {
			return []int{i2, i1}
		}

		hash[num] = i1
	}

	return nil
}
