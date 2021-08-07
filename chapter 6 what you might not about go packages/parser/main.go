package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type visitor int

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	for _, file := range os.Args[1:] {
		fmt.Println("Processing file: ", file)
		// A file set is an object that represent the source of the files
		fileSet := token.NewFileSet()
		var v visitor
		// ParseFile returns a source file parsed as a abstract syntax tree
		f, err := parser.ParseFile(fileSet, file, nil, parser.AllErrors)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// Walk walks through the abstract syntax tree `f`
		ast.Walk(v, f)
	}
}
