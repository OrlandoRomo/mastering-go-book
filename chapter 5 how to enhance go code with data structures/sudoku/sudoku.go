package main

import "fmt"

func ValidPuzzle(board [][]int) bool {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			// iEl holds every submatrix row index
			iEl := i * 3
			// jEl hold every submatrix column index
			jEl := j * 3
			cellsChecked := make([]int, 9)
			for k := 0; k <= 2; k++ {
				for m := 0; m <= 2; m++ {
					// iBig holds the row position of the current element of th submatrix
					iBig := iEl + k
					// jBig holds the column position of the current element of th submatrix
					jBig := jEl + m
					val := board[iBig][jBig]
					if val > 0 && val < 10 {
						if cellsChecked[val-1] == 1 {
							fmt.Printf("value: %d appear two times\n", val)
							return false
						}
						cellsChecked[val-1] = 1
					}
				}
			}
		}
	}
	// Rows
	for i := 0; i <= 8; i++ {
		sum := 0
		for j := 0; j <= 8; j++ {
			sum += board[i][j]
		}
		if sum != 45 {
			return false
		}
	}

	// Columns
	for i := 0; i <= 8; i++ {
		sum := 0
		for j := 0; j <= 8; j++ {
			sum += board[j][i]
		}
		if sum != 45 {
			return false
		}
	}
	return true
}
