package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some fatal error")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.SetOutput(sysLog)
	log.Fatal(sysLog)
	fmt.Println("You're seening this")
}
