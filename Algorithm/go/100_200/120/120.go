package main

/*
思路:
dp[i][j]: 从顶点到当前节点的最小路径和
注意：
1、需要先处理边界
*/
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	var result int
	for i := 0; i < m; i++ {
		for j := 0; j <= i; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = triangle[i][j]
			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][j] + triangle[i][j]
				} else if j == i {
					dp[i][j] = dp[i-1][j-1] + triangle[i][j]
				} else {
					dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
				}
			}

			if i == m-1 {
				if j == 0 {
					result = dp[i][j]
				} else {
					if dp[i][j] < result {
						result = dp[i][j]
					}
				}
			}
		}
	}
	return result
}

func main() {

}
