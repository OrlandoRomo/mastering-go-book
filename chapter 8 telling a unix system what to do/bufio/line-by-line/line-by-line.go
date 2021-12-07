package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type FileSet struct {
	Names []string
}

func (f *FileSet) GetFileNames() []string {
	return f.Names
}

func (f *FileSet) String() string {
	return fmt.Sprint(f.Names)
}

func (f *FileSet) Set(s string) error {
	if len(f.Names) > 0 {
		return errors.New("Cannot use name flags more the once!")
	}

	names := strings.Split(s, ",")
	for _, name := range names {
		f.Names = append(f.Names, name)
	}

	return nil
}

func main() {
	fileSet := FileSet{}
	flag.Var(&fileSet, "files", "file names separated by comma")

	flag.Parse()

	for _, file := range fileSet.GetFileNames() {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Print(line)
	}
	return nil
}
