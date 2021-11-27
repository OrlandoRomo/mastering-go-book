package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type NameFlag struct {
	Names []string
}

func (n *NameFlag) GetNames() []string {
	return n.Names
}

/*Implement the methods of the flag.Value interface
type Value interface {
	String() string
	Set(string) error
}*/

func (n *NameFlag) String() string {
	return fmt.Sprint(n.Names)
}

func (n *NameFlag) Set(s string) error {
	if len(n.Names) > 0 {
		return errors.New("Cannot use name flags more the once!")
	}
	names := strings.Split(s, ",")
	for _, name := range names {
		n.Names = append(n.Names, name)
	}

	return nil
}

func main() {

	var manyNames NameFlag
	minusK := flag.Int("K", 0, "K is an integer")
	minusO := flag.String("O", "Orlando", "O is name")
	flag.Var(&manyNames, "names", "comma-separated list")

	flag.Parse()

	fmt.Println("-K", *minusK)
	fmt.Println("-O", *minusO)

	for i, name := range manyNames.GetNames() {
		fmt.Println(i, name)
	}
	fmt.Println("Remaining command line arguments: ")
	for i, arg := range flag.Args() {
		fmt.Println(i, ":", arg)
	}

}
