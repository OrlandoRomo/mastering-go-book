package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
)

var SIZE = 2
var GLOBAL = 0
var LOCAL = 0

type Visitor struct {
	// ast.GenDecl represents any general declarations such as variables, imports, contants
	Package map[*ast.GenDecl]bool
}

func (v Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	switch d := n.(type) {
	// short variable declaration
	case *ast.AssignStmt:
		// token.DEFINE means :=
		if d.Tok != token.DEFINE {
			return v
		}
		for _, name := range d.Lhs {
			v.IsItLocal(name)
		}
	// Statement with a range clause
	case *ast.RangeStmt:
		v.IsItLocal(d.Key)
		v.IsItLocal(d.Value)
	// function declaration
	case *ast.FuncDecl:
		// if d as a func declaration contains receivers func(v Receiver) Name(){}
		if d.Recv != nil {
			v.CheckAll(d.Recv.List)
		}
		// Check the required parameters of d as func declaration
		v.CheckAll(d.Type.Params.List)
		if d.Type.Results != nil {
			// Check if the function declartion returns one or more values (results)
			v.CheckAll(d.Type.Results.List)
		}
	// General declaration represents imports, constants, type and variables declarations
	case *ast.GenDecl:
		// token.VAR represents the keyword VAR
		if d.Tok == token.VAR {
			return v
		}
		// The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.
		for _, spec := range d.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if ok {
				for _, name := range valueSpec.Names {
					nameLength := len(name.Name)
					// if the declaration starts with an undercore _ continue with the next constant or variable declaration
					if name.Name == "_" {
						continue
					}
					// if the constant or variable declaration exists in the package map it means is a global variable
					if v.Package[d] {
						if nameLength == SIZE {
							fmt.Printf("* %s\n", name.Name)
							GLOBAL++
							continue
						}
					}
					// The constant or variable declaration is a local variable
					if nameLength == SIZE {
						fmt.Printf("* %s\n", name.Name)
						LOCAL++
					}
				}
			}
		}
	}
	return v
}

func (v Visitor) IsItLocal(n ast.Node) {
	identifier, ok := n.(*ast.Ident)
	if !ok {
		return
	}
	if identifier.Name == "_" || identifier.Name == "" {
		return
	}
	if identifier.Obj != nil && identifier.Obj.Pos() == identifier.Pos() {
		if len(identifier.Name) == SIZE {
			fmt.Printf("* %s\n", identifier.Name)
			LOCAL++
		}
	}
}

// A Field represents a Field declaration list in a struct type,
// a method list in an interface type, or a
// parameter/result declaration in a signature.
func (v Visitor) CheckAll(fs []*ast.Field) {
	for _, f := range fs {
		for _, name := range f.Names {
			v.IsItLocal(name)
		}
	}
}

func MakeVisitor(f *ast.File) Visitor {
	pkg := make(map[*ast.GenDecl]bool, 0)
	for _, declaration := range f.Decls {
		dl, ok := declaration.(*ast.GenDecl)
		if ok {
			pkg[dl] = true
		}
	}
	return Visitor{
		Package: pkg,
	}
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("Not enough arguments")
		return
	}
	temp, err := strconv.Atoi(os.Args[1])
	if err != nil {
		SIZE = 2
		fmt.Println("Using default SIZE: ", SIZE)
	}
	if err == nil {
		SIZE = temp
	}
	var v Visitor
	fileSet := token.NewFileSet()
	for _, file := range os.Args[2:] {
		fmt.Println("Processing file: ", file)
		f, err := parser.ParseFile(fileSet, file, nil, parser.AllErrors)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		v = MakeVisitor(f)
		ast.Walk(v, f)
	}
	fmt.Printf("Local: %d\n", LOCAL)
	fmt.Printf("Global: %d\n", GLOBAL)

}
