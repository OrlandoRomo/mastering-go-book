package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.SetOutput(sysLog)
	log.Println("LOG_INFO + LOCAL_7: Logging in go")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "some program!")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.SetOutput(sysLog)
	log.Println("LOG_MAIL: Loggin in Go!")
	fmt.Println("Hello my gorgeous students")

}
