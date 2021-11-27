package main

import (
	"fmt"
	"os"
	"reflect"
)

type A struct {
	X int
	Y float64
	Z string
}

type B struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	//ValueOf returns an initialized reflection.Value of the concrete value of i
	//Elem() function returns the value that the interface is storing
	xReflect := reflect.ValueOf(&x).Elem()
	// Type() function returns the concrete Go type of reflect.Value
	xType := xReflect.Type()
	fmt.Printf("The type x is %s\n", xType)

	a := A{100, 200.12, "Struct A"}
	b := B{1, 2, "Struct B", -1.2}
	var r reflect.Value

	arguments := os.Args
	if len(arguments) == 1 {
		r = reflect.ValueOf(&a).Elem()
	} else {
		r = reflect.ValueOf(&b).Elem()
	}

	xType = r.Type()
	// main.A or main.B
	fmt.Printf("i Type: %s\n", xType)
	// NumField returns the number of fields of the structure reflection.Value
	fmt.Printf("The %d fields of %s are: \n", r.NumField(), xType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: %v ", xType.Field(i).Name)
		fmt.Printf("with type: %v ", xType.Field(i).Type)
		fmt.Printf("and value: %v \n", r.Field(i).Interface())
	}
}
