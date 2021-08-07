package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	for _, file := range os.Args[1:] {
		fmt.Println("Reading file: ", file)
		// Read file from arguments
		f, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// A file set is an object that represent the source of the files
		fileSet := token.NewFileSet()
		files := fileSet.AddFile(file, fileSet.Base(), len(f))
		var scan scanner.Scanner
		// Prepare the scanner to tokenize the source file and take every token as a comment
		scan.Init(files, f, nil, scanner.ScanComments)

		for {
			pos, tok, lit := scan.Scan()
			if tok == token.EOF {
				break
			}
			fmt.Printf("%s\t%s\t%q\n", fileSet.Position(pos), tok, lit)
		}
	}

}
