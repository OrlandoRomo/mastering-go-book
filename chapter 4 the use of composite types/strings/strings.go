package main

import "fmt"

func main() {
	const strLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	emoji := "👀🤡"
	fmt.Printf("emoji length %d\n", len(emoji))
	fmt.Printf("% x\n", emoji)
	fmt.Println(strLiteral)
	// %x will return the AB part of the \xAB
	fmt.Printf("x: %x\n", strLiteral)
	fmt.Printf("strLiteral length: %d\n", len(strLiteral))

	for i := 0; i < len(strLiteral); i++ {
		// %x prints the bytes
		fmt.Printf("%x ", strLiteral[i])
	}
	fmt.Println()

	// %q will print in double-quoate string
	fmt.Printf("q: %q\n", strLiteral)
	// %+q will print in double-quoate string
	fmt.Printf("+q: %+q\n", strLiteral)
	// % x will print the bytes separated by a whitespace
	fmt.Printf("q: % x\n", strLiteral)

	fmt.Printf("s: As a string: %s\n", strLiteral)

	s2 := "å∑∫∂"
	for x, y := range s2 {
		// %#U print the character in format U-0058
		fmt.Printf("%#U starts at byt position %d\n", y, x)
	}

	const s3 = "ab12AB"
	fmt.Println("s3: ", s3)
	fmt.Printf("x: % x\n", s3)

	fmt.Printf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
