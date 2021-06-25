package main

import (
	"errors"
	"fmt"
)

func ManyValuesAndError() (int, int, int, int, error) {
	return 1, 2, 3, 4, errors.New("An error here")
}

func main() {
	// here we only care about the err returned by the ManyValuesAndError function
	_, _, _, _, err := ManyValuesAndError()
	if err != nil {
		fmt.Println(err)
	}
	n1, n2, n3, n4, _ := ManyValuesAndError()
	fmt.Println("n1", n1)
	fmt.Println("n2", n2)
	fmt.Println("n3", n3)
	fmt.Println("n4", n4)
}
