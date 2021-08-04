package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("missing arguments")
		os.Exit(1)
	}
	file := arguments[1]
	board, err := ImportFile(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	if ValidPuzzle(board) {
		fmt.Println("Yes, the puzzle is valid")
	}
}
