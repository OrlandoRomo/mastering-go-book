package main

import (
	"io"
	"os"
)

func main() {
	arguments := os.Args
	myString := ""
	if len(arguments) == 1 {
		myString = "Please provide an argument"
	}
	if len(arguments) > 1 {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, "This is the standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
