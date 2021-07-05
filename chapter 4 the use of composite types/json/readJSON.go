package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
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
	now := time.Now()
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Provide a JSON file")
		os.Exit(1)
	}
	filename := arguments[1]
	var record Record
	err := loadJSON(filename, &record)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(record)
	duration := time.Since(now)
	fmt.Println(duration)

}

func loadJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()
	decoder := json.NewDecoder(in)
	err = decoder.Decode(key)
	if err != nil {
		return err
	}
	return nil
}
