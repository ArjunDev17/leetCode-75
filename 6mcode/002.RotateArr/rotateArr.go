package main

import "fmt"

func main() {

	// Input array
	arr := []int{1, 2, 3, 4}

	// Number of left rotations
	k := 2

	n := len(arr)

	// Handle case when k > n
	k = k % n

	// Step 1: Reverse first k elements
	reverse(arr, 0, k-1)

	// Step 2: Reverse remaining elements
	reverse(arr, k, n-1)

	// Step 3: Reverse entire array
	reverse(arr, 0, n-1)

	// Print result
	fmt.Println(arr)
}

// reverse reverses elements of the slice between start and end index
func reverse(arr []int, start int, end int) {

	for start < end {

		// Swap elements
		arr[start], arr[end] = arr[end], arr[start]

		// Move pointers
		start++
		end--
	}
}
