package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 1 // Start from the second element
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}

	return i // 'i' will be the number of unique elements
}
