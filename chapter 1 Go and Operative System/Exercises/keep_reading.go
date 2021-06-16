// Write a program that keeps reading the user input until it gets the word of END
// to finish the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func showInput() {
	fmt.Print("Give me an integer: ")
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	showInput()
	for scanner.Scan() {
		argument := scanner.Text()
		if strings.ToLower(argument) == "end" {
			os.Exit(1)
		}
		_, err := strconv.ParseFloat(argument, 16)
		if err != nil {
			fmt.Println("The provied argument that is not an integer")
			showInput()
			continue
		}
		fmt.Printf("You typed: %v\n", argument)
		showInput()
	}

}
