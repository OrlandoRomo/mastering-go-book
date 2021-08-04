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
	if len(arguments) != 5 {
		fmt.Println("Wrong number of arguments")
		return
	}
	var rowM1, colM1, rowM2, colM2 int
	rowM1, colM1, err := parseDimensions(arguments[1], arguments[2])
	if err != nil {
		fmt.Println(err.Error())
	}
	rowM2, colM2, err = parseDimensions(arguments[3], arguments[4])
	if err != nil {
		fmt.Println(err.Error())
	}

	if colM1 <= 0 || rowM1 <= 0 || colM2 <= 0 || rowM2 <= 0 {
		fmt.Println("Need positive dimensions")
		return
	}
	if colM1 != rowM2 {
		fmt.Printf("Could not perform array multiplication using n=%d,%d\n", colM1, rowM2)
		return
	}
	fmt.Printf("Using %dx%d and %dx%d arrays\n", rowM1, colM1, rowM2, colM2)
	m1 := make([][]int, rowM1)
	m2 := make([][]int, rowM2)

	createValues(m1, rowM1, colM1)
	createValues(m2, rowM2, colM2)

	fmt.Println("M1: ", m1)
	fmt.Println("M2: ", m2)
	now := time.Now()
	fmt.Println(now)
	result := multiplyMatrices(m1, m2, rowM1, colM2)
	diff := time.Since(now)
	fmt.Println(diff)
	fmt.Println("Result: ", result)
}

func random(min, max int) int {
	return rand.Intn(max-min) + max
}

func multiplyMatrices(m1, m2 [][]int, rowM1, colM2 int) [][]int {
	result := make([][]int, rowM1)
	for i := 0; i < rowM1; i++ {
		result[i] = make([]int, colM2)
		for j := 0; j < colM2; j++ {
			for k := 0; k < len(m2); k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return result
}

func parseDimensions(s1, s2 string) (int, int, error) {
	d1, err := strconv.Atoi(s1)
	if err != nil {
		return 0, 0, err
	}
	d2, err := strconv.Atoi(s2)
	if err != nil {
		return 0, 0, nil
	}
	return d1, d2, nil
}

func createValues(m [][]int, row, column int) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			m[i] = append(m[i], random(-5, i*j))
		}
	}
}
