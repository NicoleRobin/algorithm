package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	var maxVal int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = 0

			} else {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i][j-1])) + 1
				}
			}
			if dp[i][j] > maxVal {
				maxVal = dp[i][j]
			}
		}
	}
	return maxVal
}

func main() {
	testCases := []struct {
		matrix [][]byte
		expect int
	}{
		{
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
			},
			expect: 1,
		},
	}

	for _, testCase := range testCases {
		result := maximalSquare(testCase.matrix)
		fmt.Println("result:", result, " expect:", testCase.expect)
	}
}
