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
	fmt.Println(now)
	recordsJSON := `
	{
		"Name": "Yeah",
		"Surname": "I like this 🥳",
		"Phones": [
			{
				"Mobile": true,
				"Number": "4493429422"
			},
			{
				"Mobile": true,
				"Number": "23422323"
			}
		]
	}
	`
	var record Record
	err := json.Unmarshal([]byte(recordsJSON), &record)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(record)
	duration := time.Since(now)
	fmt.Println(duration)

}
