package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var myTime string
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	myTime = os.Args[1]
        parsedDate, err := time.Parse("15:04", myTime)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Full: ", parsedDate)
	fmt.Println("Time: ", parsedDate.Hour(), parsedDate.Minute())

}
