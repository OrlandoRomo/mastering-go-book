package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

const Kb = 0.001

func main() {
	ReadFileAtOne()
	ReadChunks()
}

func ReadFileAtOne() {
	now := time.Now()
	contentBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	diff := time.Now().Sub(now)
	fmt.Println("Took: ", diff)
	fmt.Println("Length of kB of the file ", float64(len(contentBytes))*Kb)
}

func ReadChunks() {
	now := time.Now()
	nBytes, nChunks := int64(0), int64(0)
	file, err := os.Open(os.Args[1])

	if err != nil {
		return
	}
	defer file.Close()
	r := bufio.NewReader(file)
	// Creating chunks
	buffer := make([]byte, 0, 4*1024)
	for {
		n, err := r.Read(buffer[:cap(buffer)])
		if n == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				continue
			}
		}
		nChunks++
		nBytes += int64(n)
	}

	diff := time.Now().Sub(now)
	fmt.Println("Took: ", diff)
	fmt.Println("Chunks: ", nChunks)
	fmt.Println("Length of kB of the file ", float64(nBytes)*Kb)
}
