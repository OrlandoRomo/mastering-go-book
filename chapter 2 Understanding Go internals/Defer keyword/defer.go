package main

import (
	"fmt"
)

func f1() {
	// prints 1 2 3
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func f2() {
	// prints 0 0 0
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func f3() {
	//prints 1 2 3
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)

	}
}

func main() {
	f1()
	f2()
	fmt.Println()
	f3()
	fmt.Println()
}
