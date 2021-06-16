package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{0, 1, 2, 3, 4, 5}
	// Assing the memory allocation of &array[0]
	pointer := &array[1]
	// Deference the pointer
	fmt.Print(*pointer, " \n")
	// Getting the address memory of
	memAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memAddress))
		fmt.Print(*pointer, " ")
		memAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}
	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memAddress))
	fmt.Print("One more: ", *pointer, " ")
	memAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println()
}
