package main

import (
	"fmt"
	"time"
)

const (
	windowSize = 200000
	msgCount   = 1000000
)

type message []byte

type channel [windowSize]message

var worst time.Duration

func mkMessage(n int) message {
	m := make(message, 1024)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

func pushMessage(c *channel, highID int) {
	start := time.Now()
	m := mkMessage(highID)
	(*c)[highID/windowSize] = m
	elapased := time.Since(start)
	if elapased > worst {
		worst = elapased
	}
}

func main() {
	var c channel
	for i := 0; i < msgCount; i++ {
		pushMessage(&c, i)
	}
	fmt.Println("Thw worst push time: ", worst)
}
