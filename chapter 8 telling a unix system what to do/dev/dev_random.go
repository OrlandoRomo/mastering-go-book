package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"runtime"
)

const devRandom = "/dev/urandom"

func main() {
	switch runtime.GOOS {
	case "linux", "darwin":
		file, err := os.Open(devRandom)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var seed int64
		binary.Read(file, binary.LittleEndian, &seed)
		fmt.Println("seed ", seed)
	}
}
