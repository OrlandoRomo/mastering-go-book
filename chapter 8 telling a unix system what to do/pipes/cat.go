package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := ""
	arguments := os.Args
	if len(arguments) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	for i := 1; i < len(arguments); i++ {
		filename = arguments[i]
		err := printFile(filename)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func printFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}
