package main

import (
	"encoding/json"
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

var JsonFile = "save.json"

func main() {
	file, err := os.OpenFile(JsonFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	records := Record{
		Name:    "Joe",
		Surname: "Sample",
		Phones: []Phone{
			{
				Mobile: true,
				Number: "4493429422",
			},
			{
				Mobile: true,
				Number: "23422323",
			},
		},
	}
	saveJSON(file, records)
	defer file.Close()
}

func saveJSON(file *os.File, key interface{}) {
	encoder := json.NewEncoder(file)
	err := encoder.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}
