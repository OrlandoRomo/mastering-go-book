package main

import (
	"fmt"
	"runtime"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc", mem.Alloc)
	fmt.Println("mem.TotalAlloc", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NumGC)
	fmt.Println("---------")
}
func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		// Allocate a large amount of memory with every go slices
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed")
		}
	}
	// prints stats of the garbage collector
	printStats(mem)

	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			// Allocate another large amount of memory with every go slices
			s := make([]byte, 100000000)
			if s == nil {
				fmt.Println("Operation failed")
			}
		}
	}
	// prints stats of the garbage collector
	printStats(mem)
}
