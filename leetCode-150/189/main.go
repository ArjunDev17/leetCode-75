package main

import (
	"fmt"
)

// Function to rotate array to the right by k steps
func rotate(nums []int, k int) {

	// Length of array
	n := len(nums)

	// Handle cases where k > n
	k = k % n

	// Step 1: Reverse entire array
	reverse(nums, 0, n-1)

	// Step 2: Reverse first k elements
	reverse(nums, 0, k-1)

	// Step 3: Reverse remaining elements
	reverse(nums, k, n-1)
}

// Helper function to reverse a portion of array
func reverse(nums []int, start int, end int) {

	// Swap elements from start and end moving inward
	for start < end {

		// Swap nums[start] and nums[end]
		nums[start], nums[end] = nums[end], nums[start]

		// Move pointers
		start++
		end--
	}
}

func main() {

	// Example input
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	// Rotate array
	rotate(nums, k)

	// Output result
	fmt.Println("Rotated Array:", nums)
}