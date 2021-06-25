package main

import "fmt"

func main() {
	const r1 = 'å'
	fmt.Println("(int32) r1: ", r1)
	fmt.Printf("(HEX) %x\n", r1)
	fmt.Printf("(As a string) r1: %s\n", string(r1))
	fmt.Printf("(As a character) r1: %c\n", r1)

	fmt.Println("a string is collection of runes", []byte("Orlando 🤡"))
	aString := []byte("Orlando 🥺")
	for x, y := range aString {
		fmt.Println("int32: ", y)
		fmt.Printf("Char: %c\n", aString[x])
	}
	fmt.Printf("%s\n", aString)
}
