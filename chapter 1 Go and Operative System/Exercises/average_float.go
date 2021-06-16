// Write a program that finds the average value of floating-point numbers that are
// given by the command-line
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
		fmt.Println("Provide a floating-point number")
		return
	}
	var err error = errors.New("An error")
	k := 1
	for err != nil {
		if k >= argsLength {
			fmt.Println("None of the arguments is a floating-point number")
			return
		}
		_, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	average := 0.0
	counter := 0
	for i := 1; i < argsLength; i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			average += n
			counter++
		}
	}
	average /= float64(counter)
	fmt.Printf("Average float-point number: %.2f\n", average)
}
