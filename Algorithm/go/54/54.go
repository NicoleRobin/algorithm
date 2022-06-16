package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			fmt.Printf("i:%d, j:%d, num:%d\n", top, column, matrix[top][column])
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			fmt.Printf("i:%d, j:%d, num:%d\n", row, right, matrix[row][right])
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				fmt.Printf("i:%d, j:%d, num:%d\n", bottom, column, matrix[bottom][column])
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				fmt.Printf("i:%d, j:%d, num:%d\n", row, left, matrix[row][left])
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

func main() {
	fmt.Println("leetcode-54")
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := spiralOrder(matrix)
	fmt.Println(result)
}
