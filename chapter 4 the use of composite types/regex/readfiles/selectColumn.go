// to run: go run selectColumn.go 3 text1.txt index.html main.rs | head text1.txt index.html main.rs
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: selectColumn column <file1> [file2] [...fileN]\n")
		os.Exit(0)
	}
	temp, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Column value is not an integer", temp)
		return
	}
	column := temp
	if column < 0 {
		fmt.Println("Invalid column number!")
		os.Exit(0)
	}
	for _, filename := range arguments[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %a\n", err)
			continue
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("error reading file %s", err)
				break
			}
			data := strings.Fields(line)
			if len(data) > 0 {
				fmt.Println(data[column-1])
			}
		}
	}
}
