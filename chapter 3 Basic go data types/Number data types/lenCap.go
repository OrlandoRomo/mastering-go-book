package main

import (
	"fmt"
)

func main() {

	slice := make([]int, 0)
	fmt.Println("The length of the slice is ", len(slice))
	fmt.Println("The capacity of the slice is ", cap(slice))
	slice = append(slice, 1) // length = 1 and cap = 1
	slice = append(slice, 2) // length = 2 and cap = 2
	slice = append(slice, 3) // length = 3 and cap = 4
	slice = append(slice, 4) // length = 4 and cap = 4
	slice = append(slice, 5) // length = 5 and cap = 8
	slice = append(slice, 6) // length = 6 and cap = 8

	fmt.Println("The length of the slice is ", len(slice))
	fmt.Println("The capacity of the slice is ", cap(slice))
}
