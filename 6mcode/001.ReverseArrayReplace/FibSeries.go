package main

import "fmt"

func FibSeries(times int) {
	a, b := 0, 1

	for range times {
		fmt.Println(" ", a)
		b, a = a, b+a

	}

	/*
		for i := 0; i < times; i++ {
			fmt.Println(" ",a)
			b,a=a,b+1
		}
	*/
}
