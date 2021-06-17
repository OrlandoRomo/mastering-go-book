package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
  var myDate string
  if len(os.Args) != 2 {
    fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
    os.Exit(1)
  }
  myDate = os.Args[1]
  dateParsed, err:= time.Parse("02 January 2006", myDate)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  fmt.Println("Full: ", dateParsed)
  fmt.Printf("Time: %v %v %v\n", dateParsed.Day(), dateParsed.Month(), dateParsed.Year())
}
