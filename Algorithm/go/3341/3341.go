package main

import (
	"math"
)

func minTimeToReach(moveTime [][]int) int {
	m := len(moveTime)
	if m == 0 {
		return 0
	}
	n := len(moveTime[0])

	grid := make([][]int, len(moveTime))
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			left, up := math.MaxInt32, math.MaxInt32
			if i > 0 {
				left = grid[i-1][j]
			}
			if j > 0 {
				up = grid[i][j-1]
			}

			tempMin := min(left, up)
			if tempMin < moveTime[i][j] {
				tempMin = moveTime[i][j]
			}
			grid[i][j] = tempMin + 1

			// fmt.Println("i:", i, "j:", j, "grid:", grid)
		}
	}

	return grid[m-1][n-1]
}

type TestCase struct {
	MoveTime [][]int
	Expected int
}

func main() {
	testCases := []TestCase{
		{
			MoveTime: [][]int{
				{0, 4},
				{4, 4},
			},
			Expected: 6,
		},
		{
			MoveTime: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			Expected: 3,
		},
	}
	for i, testCase := range testCases {
		result := minTimeToReach(testCase.MoveTime)
		if result != testCase.Expected {
			println("Test case", i, "failed. Expected", testCase.Expected, "but got", result)
		} else {
			println("Test case", i, "passed.")
		}
	}
}
