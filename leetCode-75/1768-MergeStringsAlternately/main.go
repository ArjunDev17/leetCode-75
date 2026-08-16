package main

import "fmt"

func PrintToStringCh(str, str1 string) {
	fmt.Println("str ", str+"\nstr1 :", str1)
}

// Print Characters Alternately
func PrintToStringCharAlterNative(str, str1 string) {
	a, b := 0, 0
	for i := 0; i < len(str) || i < len(str1); i++ {
		if len(str)-1 >= a {
			fmt.Printf(" %c ", str[i])
			a++
		}
		if len(str1) >= b {
			fmt.Printf(" %c ", str1[i])
			b++
		}

	}
}
func mergeAlternately(word1 string, word2 string) string {
	i, j := 0, 0

	result := make([]byte, 0, len(word1)+len(word2))

	for i < len(word1) || j < len(word2) {

		// Take character from word1
		if i < len(word1) {
			result = append(result, word1[i])
			i++
		}

		// Take character from word2
		if j < len(word2) {
			result = append(result, word2[j])
			j++
		}
	}

	return string(result)
}

func addMinimum1(word string) int {
	// 	target:="abc"
	// 	h1:="a"
	// 	h2:="b"
	// 	if word==target{
	// 		return 0
	// 	}
	// 	if word==""

	return 0

}

func addMinimum(word string) int {
	additions := 0
	expected := byte('a')

	for i := 0; i < len(word); i++ {
		ch := word[i]

		// Insert missing characters until we reach ch
		for expected != ch {
			additions++

			if expected == 'a' {
				expected = 'b'
			} else if expected == 'b' {
				expected = 'c'
			} else {
				expected = 'a'
			}
		}

		// Current character is matched
		// Move to the next expected character
		if expected == 'a' {
			expected = 'b'
		} else if expected == 'b' {
			expected = 'c'
		} else {
			expected = 'a'
		}
	}

	// Complete the last "abc" group
	for expected != 'a' {
		additions++

		if expected == 'b' {
			expected = 'c'
		} else {
			expected = 'a'
		}
	}

	return additions
}

func main() {
	PrintToStringCharAlterNative("ab", "xyz")
}
