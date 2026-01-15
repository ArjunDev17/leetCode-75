package main

import (
	"fmt"
)

func main() {
	fmt.Println("Reverse array in-place")

	reverseArr()

	// FibSeries(8)
}

func reverseArr() {

	arr := []int{3, 4, 11, 7, 9}

	left := 0
	right := len(arr) - 1

	for left < right {

		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	fmt.Println(arr)

}
