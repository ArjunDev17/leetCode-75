package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	// k will track the position where next valid element should go
	k := 0

	// Loop through all elements
	for i := 0; i < len(nums); i++ {

		// If current element is NOT equal to val
		if nums[i] != val {

			// Place it at index k
			nums[k] = nums[i]

			// Move k forward
			k++
		}
	}

	// k is the count of elements not equal to val
	return k
}

func main() {
	// Example input
	nums := []int{3, 2, 2, 3}
	val := 3

	// Call function
	k := removeElement(nums, val)

	// Output result
	fmt.Println("k:", k)
	fmt.Println("Updated nums:", nums[:k]) // Only print valid part
}