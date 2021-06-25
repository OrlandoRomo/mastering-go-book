package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name     string
	Age      int
	Birthday time.Time
	Hobbies  []string
}

func main() {
	orlando := new(Person) // return a pointer of Person, zero-value is nil
	fmt.Println(&orlando)

	pointerSlice := new([]string) // *[]string, zero-value is nil
	fmt.Println(&pointerSlice)

	allocatedSlice := make([]string, 0)
	fmt.Println(allocatedSlice)
}
