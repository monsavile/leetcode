package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		complement := target - num

		if j, ok := m[complement]; ok {
			return []int{i, j}
		}

		m[num] = i
	}

	return []int{}
}
