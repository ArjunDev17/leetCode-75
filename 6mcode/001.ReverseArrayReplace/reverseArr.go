package main

import "fmt"

func main() {
	fmt.Println("Reverse array in-place")

	reverseArr()
}

func reverseArr() {

	// Input array
	arr := []int{2, 4, 11, 33, 7}

	// Left and right pointers
	left := 0
	right := len(arr) - 1

	// Reverse in-place
	for left < right {

		// Swap elements
		arr[left], arr[right] = arr[right], arr[left]

		// Move pointers
		left++
		right--
	}

	// Output result
	fmt.Println(arr)
}
