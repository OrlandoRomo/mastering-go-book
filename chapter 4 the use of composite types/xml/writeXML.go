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

var XMLFile = "save.xml"

func main() {
	file, err := os.OpenFile(XMLFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	records := Record{
		Name:    "The crusaders",
		Surname: "I love you Dua Lipa",
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
	encoder := xml.NewEncoder(file)
	err := encoder.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}
