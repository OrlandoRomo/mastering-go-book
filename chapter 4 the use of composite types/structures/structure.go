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

// We could set default values by given the arguments of the arguments of
// the function or simply return the pointer to the structure `Person`
func New(name string) *Person {
	return &Person{
		Name: name,
	}
}

// Or better
func NewPerson() *Person {
	return &Person{}
}

func main() {
	orlando := Person{
		Age:  23,
		Name: "Orlando",
	}

	fmt.Println("This is the name", orlando.Name)
	fmt.Println("This is the age", orlando.Age)

	another := New("Dua Lipa")
	fmt.Println("Name", another.Name)
}
