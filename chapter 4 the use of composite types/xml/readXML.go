package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Phones  []Phone
}

type Phone struct {
	Mobile bool
	Number string
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Provide a JSON file")
		os.Exit(1)
	}
	filename := arguments[1]
	var record Record
	err := loadXML(filename, &record)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(record)

}

func loadXML(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()
	decoder := xml.NewDecoder(in)
	err = decoder.Decode(key)
	if err != nil {
		return err
	}
	return nil
}
