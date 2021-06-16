package main

import (
	"fmt"
	"log"
	"os"
)

// Path of the log file
var LOGFILE = "/tmp/mGo.log"

func main() {
	// 644 means rw-r--r---
	// User can read and write
	// Group only can read
	// Other only can read
	f, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// Creating the logger
	iLog := log.New(f, "customeLineNumber: ", log.LstdFlags)

	// log.LstdFlags has the flags of time and date of each log entry
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
