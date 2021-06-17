package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := make([]int, 100)
	for i, _ := range numbers {
		numbers[i] = rand.Intn(400)
	}
	mapNumbers := make(map[int]string)
	for _, num := range numbers {
		mapNumbers[num] = time.Now().String()
	}
	for key, num := range mapNumbers {
		fmt.Printf("key: %v \t value: %v\n", key, num)
	}
}
