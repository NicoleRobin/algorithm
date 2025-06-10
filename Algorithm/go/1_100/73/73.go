package main

import "fmt"

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	mLen := len(matrix)
	nLen := len(matrix[0])

	zeroRows := make(map[int]bool)
	zeroCols := make(map[int]bool)
	for i := 0; i < mLen; i++ {
		for j := 0; j < nLen; j++ {
			if matrix[i][j] == 0 {
				zeroRows[i] = true
				zeroCols[j] = true
			}
		}
	}
	for row := range zeroRows {
		for j := 0; j < nLen; j++ {
			matrix[row][j] = 0
		}
	}
	for col := range zeroCols {
		for i := 0; i < mLen; i++ {
			matrix[i][col] = 0
		}
	}
}

func main() {
	fmt.Println("leetcode 73")

	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)

	for _, row := range matrix {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
