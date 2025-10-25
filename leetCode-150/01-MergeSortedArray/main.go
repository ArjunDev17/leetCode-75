package main

// Brute-forece approche 
func mergeSortedArr(num1 []int, num2 []int, m, n int) {
	for i := 0; i < len(num1); i++ {
		ele := num2[i]

		j := m + i - 1

		if j >= 0 && num1[j] > ele {
			num1[j+1] = num1[j]
			j--
		}
		num1[j] = ele
	}
}
