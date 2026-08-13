package main

import "fmt"

func PrintToStringCh(str, str1 string) {
	fmt.Println("str ", str+"\nstr1 :", str1)
}

// Print Characters Alternately
func PrintToStringCharAlterNative(str, str1 string) {
	for i := 0; i < len(str) || i < len(str1); i++ {
		if len(str) > 0 {
			fmt.Printf(" %c ", str[i])
		}
		if len(str1) > 0 {
			fmt.Printf(" %c ", str1[i])
		}

	}
}

func main() {
	PrintToStringCharAlterNative("Ram", "Sheeta")
	//R S a h m e
}
