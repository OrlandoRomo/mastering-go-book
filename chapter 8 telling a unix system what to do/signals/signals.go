package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// creating buffer chanel of os.Signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINFO)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught: ", sig)
			case syscall.SIGINT:
				handlingSignal(sig)
				return
			}
		}
	}()

	for {
		fmt.Println(".")
		time.Sleep(20 * time.Second)
	}

}

func handlingSignal(s os.Signal) {
	fmt.Println("handling signal caught: ", s)
}
