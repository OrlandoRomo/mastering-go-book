package main

import (
	"encoding/xml"
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
	records := Record{
		Name:    "Dua Lipa BB",
		Surname: "🥰",
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
	xmlOut, err := xml.MarshalIndent(records, "", "	")
	xmlData := []byte(xml.Header + string(xmlOut))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(xmlData))
	duration := time.Since(now)
	fmt.Println(duration)

}
