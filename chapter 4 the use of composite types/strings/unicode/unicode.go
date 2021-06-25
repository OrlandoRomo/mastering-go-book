package main

import (
	"fmt"
	"unicode"
)

func main() {
	const str = "🥺🦀🎹🛵 check this out prro      "
	for i := 0; i < len(str); i++ {
		if unicode.IsPrint(rune(str[i])) {
			fmt.Printf("%c\n", str[i])
			continue
		}
		fmt.Println("No printable")
	}
}
