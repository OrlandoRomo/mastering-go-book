package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Wrong number of arguments")
		return
	}
	var row, col int
	row, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	col, err = strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Using %dx%d arrays\n", row, col)
	if col <= 0 || row <= 0 {
		fmt.Println("Need positive dimensions")
		return
	}
	m1 := make([][]int, row)
	m2 := make([][]int, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m1[i] = append(m1[i], random(-1, i*j+rand.Intn(10)))
			m2[i] = append(m2[i], random(-1, i*j+rand.Intn(10)))
		}
	}

	fmt.Println("M1: ", m1)
	fmt.Println("M2: ", m2)

	result := addMatrices(m1, m2)
	fmt.Println("Result: ", result)
}

func random(min, max int) int {
	return rand.Intn(max-min) + max
}

// antisymmetricallMatrix returns the opposite integer of each position of matrix
func antisymmetricallMatrix(m [][]int) [][]int {
	for i, x := range m {
		for j := range x {
			m[i][j] = -1 * m[i][j]
		}
	}
	return m
}

func addMatrices(m1, m2 [][]int) [][]int {
	result := make([][]int, len(m1))
	for i, x := range m1 {
		for j := range x {
			result[i] = append(result[i], m1[i][j]+m2[i][j])
		}
	}
	return result
}
