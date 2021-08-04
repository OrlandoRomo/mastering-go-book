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
	if rowM2 != colM2 {
		fmt.Println("The matrix B should be a square matrix")
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
	det := getDeterminant(m2, rowM2)
	if det == 0 {
		fmt.Println("The matrix B has no inverse matrix and could not perform division operation")
		return
	}
	transpose := getMatrixTranspose(m2)
	adjoint := getAdjointMatrix(transpose)
	inverse := getInverseMatrix(adjoint, det)
	divide := multiplyMatrices(m1, inverse, rowM1, colM2)
	diff := time.Since(now)
	fmt.Println(diff)
	fmt.Println(divide)
}

// random returns a random number from [min, max) range
func random(min, max int) int {
	return rand.Intn(max-min) + max
}

// parseDimensions returns the rows and columns based of s1, s2
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

// createValues returns a matrix filled up with random values
func createValues(m [][]int, row, column int) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			m[i] = append(m[i], random(-5, i*j))
		}
	}
}

// getDeterminant returns a value which represents weather the matrix m has a inverse denoted as m^-1
func getDeterminant(m [][]int, size int) int {

	if size == 1 {
		return m[0][0]
	}
	tmp := make([][]int, size-1)
	sign, det := 1, 0
	for i := 0; i < size; i++ {
		getCofactor(m, tmp, 0, i, size)
		// follwing the rule -+(1) * m[row][column] * det(Mij)
		det += sign * m[0][i] * getDeterminant(tmp, size-1)
		sign = -sign
	}
	return det
}

// getCofactor returns a matrix called cofactor avoiding row p and column q
func getCofactor(m, tmp [][]int, p, q, size int) {
	i, j := 0, 0
	for row := 0; row < size; row++ {
		if row < size-1 {
			tmp[row] = make([]int, size-1)
		}
		for column := 0; column < size; column++ {
			if row != p && column != q {
				tmp[i][j] = m[row][column]
				j++
				if j == size-1 {
					j = 0
					i++
				}
			}
		}
	}
}

// getMatrixTranspose creates a transpose matrix of m denoted as m^T
func getMatrixTranspose(m [][]int) [][]int {
	transpose := make([][]int, len(m[0]))
	for i := 0; i < len(transpose); i++ {
		transpose[i] = make([]int, len(m))
		for j := 0; j < len(transpose[i]); j++ {
			transpose[i][j] = m[j][i]
		}
	}
	return transpose
}

// getAdjointMatrix creates an adjoint matrix donted as Adj(m)
func getAdjointMatrix(m [][]int) [][]int {
	l, sign := len(m), 1
	adjoint := make([][]int, l)
	for i := range adjoint {
		adjoint[i] = make([]int, l)
	}
	tmp := make([][]int, l-1)
	if l == 1 {
		adjoint[0][0] = m[0][0]
		return adjoint
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			getCofactor(m, tmp, i, j, l)
			if (i+j)%2 == 0 {
				sign = 1
			}
			if (i+j)%2 != 0 {
				sign = -1
			}
			adjoint[j][i] = sign * getDeterminant(tmp, l-1)
			sign = -sign
		}
	}
	return adjoint
}

// getInverseMatrix returns the inverse of a matrix A denoted as A^-1
func getInverseMatrix(adjoint [][]int, det int) [][]float64 {
	length := len(adjoint)
	inverse := make([][]float64, length)
	for i := range inverse {
		inverse[i] = make([]float64, length)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			inverse[i][j] = float64(adjoint[i][j]) / float64(det)
		}
	}
	return inverse
}

func multiplyMatrices(m1 [][]int, m2 [][]float64, rowM1, colM2 int) [][]float64 {
	result := make([][]float64, rowM1)
	for i := 0; i < rowM1; i++ {
		result[i] = make([]float64, colM2)
		for j := 0; j < colM2; j++ {
			for k := 0; k < len(m2); k++ {
				result[i][j] += float64(m1[i][k]) * m2[k][j]
			}
		}
	}
	return result
}
