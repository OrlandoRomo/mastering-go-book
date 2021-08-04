package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func ImportFile(file string) ([][]int, error) {
	var cells [][]int
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		fields := strings.Fields(line)
		tmp := make([]int, 0)
		for _, field := range fields {
			value, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, value)
		}
		if len(tmp) != 0 {
			cells = append(cells, tmp)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return cells, nil
}
