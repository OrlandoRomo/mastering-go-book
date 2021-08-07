package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
)

func main() {
	varCounter, keyword := 0, "var"
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	for _, file := range os.Args {
		fmt.Println("Proccessing: ", file)
		f, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// A file set is an object that represent the source of the files
		fileSet := token.NewFileSet()
		files := fileSet.AddFile(file, fileSet.Base(), len(f))
		// Prepare the scanner to tokenize the source file and take every token as a comment
		var scan scanner.Scanner
		scan.Init(files, f, nil, scanner.ScanComments)
		for {
			_, tok, _ := scan.Scan()
			if tok == token.EOF {
				break
			}
			if tok == token.VAR {
				varCounter++
			}
		}
		fmt.Printf("The keyword %s appers %d times.\n", keyword, varCounter)
	}
}
