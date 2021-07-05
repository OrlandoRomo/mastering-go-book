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
	fmt.Println(now)
	recordsXML := `
	<?xml version="1.0" encoding="UTF-8"?>
	<Record>
			<Name>Fourplay</Name>
			<Surname>🥰</Surname>
			<Phones>
					<Mobile>true</Mobile>
					<Number>498279232</Number>
			</Phones>
			<Phones>
					<Mobile>true</Mobile>
					<Number>3123908123</Number>
			</Phones>
	</Record>
	`
	var record Record
	err := xml.Unmarshal([]byte(recordsXML), &record)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(record)
	duration := time.Since(now)
	fmt.Println(duration)

}
