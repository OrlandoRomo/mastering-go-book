// Write a program that finds the sum of all command-line arguments that are valid numbers
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	argsLength := len(arguments)
	if argsLength == 1 {
		fmt.Println("Provide an argument.")
		return
	}
	var err error = errors.New("an error")
	k := 1
	for err != nil {
		if k >= argsLength {
			fmt.Println("None of the arguments is a valid number")
			return
		}
		_, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	sum := 0
	for i := 1; i < argsLength; i++ {
		n, err := strconv.ParseFloat(arguments[i], 16)
		if err == nil {
			sum += int(n)
		}
	}
	fmt.Printf("Sum: %v\n", sum)
}
