package main

func twoSum1(nums []int, target int) []int {
	for i := range len(nums) - 1 {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		m[num] = i
	}

	for i, num := range nums {
		complement := target - num

		if j, ok := m[complement]; ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}

func twoSum3(nums []int, target int) []int {
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
