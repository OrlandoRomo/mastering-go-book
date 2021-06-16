package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5
	var p1 = &value
	// Create a pointer of int32 that points to p1 or the allocation of value
	// Any go pointer can be converted to unsafe.Pointer
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)

	*p1 = 5432123412341234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
	*p1 = 54321234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}
