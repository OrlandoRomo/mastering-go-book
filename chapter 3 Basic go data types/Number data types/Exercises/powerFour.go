package main

import "fmt"

type Power4 int

const (
	p2_0 Power4 = 1 << iota
	_
	_
	_
	p2_4
	_
	_
	_
	p2_8
	_
	_
	_
	p2_12
)

func main() {

	fmt.Println("2^0", p2_0)
	fmt.Println("2^4", p2_4)
	fmt.Println("2^8", p2_8)
	fmt.Println("2^12", p2_12)
}
