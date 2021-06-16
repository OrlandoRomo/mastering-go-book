package main

import "C"
import "sort"

//export SortArray
func SortArray(array []int) {
	sort.Ints(array)
}

func main() {}
