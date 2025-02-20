package main

import "fmt"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first, second := int(^uint(0)>>1), int(^uint(0)>>1)
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("Jai")
	nums := []int{1, 2, 3, 4, 5}
	isRes := increasingTriplet(nums)
	fmt.Println("result is :", isRes)
}
