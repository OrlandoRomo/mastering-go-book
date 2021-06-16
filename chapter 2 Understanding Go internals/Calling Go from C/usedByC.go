package main

import "C"

import "fmt"

//export PrintMessage
func PrintMessage() {
	fmt.Println("A go function")
}

//export Multiply
func Multiply(x, y int) int {
	return x * y
}

func main() {
}
