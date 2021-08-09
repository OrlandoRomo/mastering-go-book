package main

import "fmt"

func main() {
	var myInt interface{} = 123
	valInt, ok := myInt.(int)
	if ok {
		fmt.Println("Success type assertion: ", valInt)
	}
	valFloat, ok := myInt.(float64)
	if ok {
		fmt.Println(valFloat)
	} else {
		fmt.Println("Failed without panicking!")
	}

	successInt := myInt.(int)
	fmt.Println(successInt)

	// failing
	failInt := myInt.(bool)
	fmt.Println(failInt)
}
