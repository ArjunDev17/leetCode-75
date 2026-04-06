package main

import (
	"fmt"
)

func main() {

	// nums1 has extra space (size = m + n)
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3 // number of valid elements in nums1

	nums2 := []int{2, 5, 6}
	n := 3 // number of elements in nums2

	// Call merge function
	merge(nums1, m, nums2, n)

	// Print final result
	fmt.Println("Merged Array:", nums1)
}

// merge merges nums2 into nums1 in-place
func merge(nums1 []int, m int, nums2 []int, n int) {

	// Pointer to last valid element in nums1
	i := m - 1

	// Pointer to last element in nums2
	j := n - 1

	// Pointer to last position in nums1
	k := m + n - 1

	// Compare elements from the end
	for i >= 0 && j >= 0 {

		// If nums1 element is bigger, place it at the end
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			// Otherwise place nums2 element
			nums1[k] = nums2[j]
			j--
		}

		// Move result pointer
		k--
	}

	// If nums2 still has elements left, copy them
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}

	// No need to copy nums1 remaining elements
	// because they are already in correct position
}