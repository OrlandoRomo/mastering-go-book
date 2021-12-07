package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handleSignal(sig)
			case syscall.SIGTERM:
				handleSignal(sig)
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("Handling signal USR1")
			default:
				fmt.Println("Ignoring", sig)
			}
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(30 * time.Second)

	}
}

func handleSignal(s os.Signal) {
	fmt.Println("Received: ", s)
}
