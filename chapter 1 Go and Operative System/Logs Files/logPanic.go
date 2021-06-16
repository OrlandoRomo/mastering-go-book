package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some panic err")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.SetOutput(sysLog)
	log.Panic(sysLog)
	fmt.Println("Will you see this?")
}
