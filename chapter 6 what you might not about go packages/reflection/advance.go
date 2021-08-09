package main

import (
	"fmt"
	"os"
	"reflect"
)

type T1 int
type T2 int

type A struct {
	X    int
	Y    float64
	Text string
}

func (a A) CompareStruct(dest A) bool {
	rSource := reflect.ValueOf(&a).Elem()
	rDestination := reflect.ValueOf(&dest).Elem()

	for i := 0; i < rSource.NumField(); i++ {
		if rSource.Field(i).Interface() != rDestination.Field(i).Interface() {
			return false
		}
	}
	return true
}

func PrintMethods(i interface{}) {
	r := reflect.ValueOf(i)
	rType := r.Type()

	fmt.Printf("Type to examine is %s\n", rType)
	//NumMethod returns the number of methods implemented of the concrate type i
	for i := 0; i < r.NumMethod(); i++ {
		// reflection.Value.Method(i int).Type returns the definition of a function e.q func(int, int) err
		method := r.Method(i).Type()
		fmt.Println(rType.Method(i).Name, " --> ", method)
	}
}

func main() {
	x1 := T1(100)
	x2 := T2(100)
	fmt.Printf("Type of x1 is %s\n", reflect.TypeOf(x1))
	fmt.Printf("Type of x2 is %s\n", reflect.TypeOf(x2))

	var p struct{}
	r := reflect.New(reflect.TypeOf(p)).Elem()
	fmt.Printf("The type of r is %s\n", r.Type())

	a1 := A{1, 1.2, "ALV"}
	a2 := A{1, -2, "su puta mdare"}

	if a1.CompareStruct(a1) {
		fmt.Println("Equals")
	}
	if !a1.CompareStruct(a2) {
		fmt.Println("Not equal alv")
	}
	var f *os.File
	PrintMethods(f)

}
