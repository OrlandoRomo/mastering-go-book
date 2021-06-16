package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "please give an argument!"
	}
	if len(arguments) > 1 {
		myString = arguments[1]
	}

	// io.WriteString requires io.Writer interface which os.Stdout implements
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
