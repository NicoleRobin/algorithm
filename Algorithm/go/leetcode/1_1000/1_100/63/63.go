package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

			if i == 0 && j == 0 {
				dp[i][j] = 1
				continue
			}
			var up, left int
			if i-1 >= 0 {
				up = dp[i-1][j]
			}
			if j-1 >= 0 {
				left = dp[i][j-1]
			}

			dp[i][j] = left + up
		}
	}
	return dp[m-1][n-1]
}

func main() {

}
