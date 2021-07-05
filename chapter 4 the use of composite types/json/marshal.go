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
	json, err := json.Marshal(&records)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(json))
	duration := time.Since(now)
	fmt.Println(duration)

}
