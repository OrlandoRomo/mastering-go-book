package main

import (
	"flag"
	"fmt"
)

func main() {
	minusK := flag.Bool("K", true, "K flag")
	minusO := flag.Int("O", 1, "O flag")
	flag.Parse()

	valueK := *minusK
	valueO := *minusO
	valueO++

	fmt.Println("K", valueK)
	fmt.Println("-O", valueO)
}
