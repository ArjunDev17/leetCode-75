package bubleshort

import "fmt"

func main() {

	// Input array
	arr := []int{5, 2, 9, 1, 5, 6}

	n := len(arr)

	// Bubble Sort logic
	for i := 0; i < n-1; i++ {

		// Last i elements are already in correct position
		for j := 0; j < n-i-1; j++ {

			// Compare adjacent elements
			if arr[j] > arr[j+1] {

				// Swap if they are in wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	// Print sorted array
	fmt.Println("Sorted Array:", arr)
}
