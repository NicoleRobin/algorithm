package main

func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	copy(dp[0], matrix[0])

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			mn := dp[i-1][j]
			if j > 0 {
				mn = min(mn, dp[i-1][j-1])
			}
			if j < n-1 {
				mn = min(mn, dp[i-1][j+1])
			}
			dp[i][j] = mn + matrix[i][j]
		}
	}

	minVal := dp[n-1][0]
	for _, val := range dp[n-1] {
		if val < minVal {
			minVal = val
		}
	}
	return minVal

}

func main() {

}
