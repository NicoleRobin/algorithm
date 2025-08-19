package main

/*
思路：动态规划
dp[i][j]表示以(i,j)为右下角的正方形的边长
*/
func countSquares(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				dp[i+1][j+1] = min(min(dp[i+1][j], dp[i][j+1]), dp[i][j]) + 1
				ans += dp[i+1][j+1]
			}
		}
	}

	return ans
}

func main() {

}
