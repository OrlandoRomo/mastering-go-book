package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("You are using ", runtime.Compiler, " ")
	fmt.Println("On a ", runtime.GOARCH, " machine")
	fmt.Println("Using go version ", runtime.Version())
	fmt.Println("Number of CPUs ", runtime.NumCPU())
	fmt.Println("Number of goroutines ", runtime.NumGoroutine())
	fmt.Println("Go root ", runtime.GOROOT())
}
