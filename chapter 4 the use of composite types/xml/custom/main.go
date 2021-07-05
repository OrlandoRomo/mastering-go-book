package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Employee struct {
	XMLName   xml.Name `xml:"Employee"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Initials  string   `xml:"name>initials"`
	Height    float32  `xml:"height,omitempty"`
	Address
	Comment string `xml:",comment"`
}

type Address struct {
	City    string
	Country string
}

func main() {
	bobJames := &Employee{
		Id:        7,
		FirstName: "Bob",
		LastName:  "James",
		Initials:  "BJ",
		Height:    1.85,
		Address: Address{
			City:    "Missouri",
			Country: "United States",
		},
	}
	bobJames.Comment = "He is a jazz pianist"
	output, err := xml.MarshalIndent(bobJames, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output = []byte(xml.Header + string(output))
	os.Stdout.Write(output)
	os.Stdout.Write([]byte("\n"))

}
