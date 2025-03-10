package main

import "fmt"

func findMaxAvg(arr []int, k int) float64 {
	currentSum := 0
	for i := range k {
		currentSum += arr[i]
	}
	maxSum := currentSum
	for i := k; i < len(arr); i++ {
		currentSum = arr[i] - arr[i-k]

		// if maxSum < currentSum {
		// 	maxSum = currentSum  maximum avg is : 12.25
		// }
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return float64(maxSum) / float64(k)
}
func main() {
	arr := []int{1, 12, -5, -6, 50, 3}
	k := 4
	avgIs := findMaxAvg(arr, k)
	fmt.Println("maximum avg is :", avgIs)
}
