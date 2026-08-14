package main

import "fmt"

func PrintToStringCh(str, str1 string) {
	fmt.Println("str ", str+"\nstr1 :", str1)
}

// Print Characters Alternately
func PrintToStringCharAlterNative(str, str1 string) {
	a,b:=0,0
	for i := 0; (i < len(str) || i < len(str1)); i++ {
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

func main() {
	PrintToStringCharAlterNative("ab", "xyz")
	//R S a h m e
}
