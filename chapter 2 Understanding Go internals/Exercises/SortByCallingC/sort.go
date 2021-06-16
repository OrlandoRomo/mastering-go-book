package main

//#cgo CFLAGS: -I${SRCDIR}/sortlib
//#cgo LDFLAGS: ${SRCDIR}/sort.a
//#include <sort.h>
import "C"
import (
	"fmt"
)

func main() {
	var numbers = []C.int{5, 1, 4, 8, 2, 3}
	size := C.int(len(numbers))
	C.bubbleSort(&numbers[0], size)
	fmt.Println("Sort array calling C from Go")
	C.printArray(&numbers[0], size)
	fmt.Println("All done")

}
