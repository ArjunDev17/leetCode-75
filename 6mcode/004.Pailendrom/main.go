package main

import "fmt"

func main() {

	res := isPalindrome("amma")
	fmt.Println("res :", res)
}

func isPalindrome(s string) bool {

	// Left pointer starts from beginning
	left := 0

	// Right pointer starts from end
	right := len(s) - 1

	// Continue until pointers cross
	for left < right {

		// Skip non-alphanumeric characters from the left
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}

		// Skip non-alphanumeric characters from the right
		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		// Compare characters after converting to lowercase
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		// Move pointers inward
		left++
		right--
	}

	// If all characters match
	return true
}

// Helper function to check if character is alphanumeric
func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

// Helper function to convert uppercase to lowercase
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
