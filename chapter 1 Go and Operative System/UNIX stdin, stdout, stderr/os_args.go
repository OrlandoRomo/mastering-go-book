package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need to provided arguments")
		return
	}
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max := min

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Printf("Max: %v\n", max)
	fmt.Printf("Min: %v\n", min)
}
