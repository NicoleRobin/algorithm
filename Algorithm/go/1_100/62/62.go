package main

import "fmt"

/*
动态规划
1、状态定义：dp[i][j]：i,j这个位置到出口的路线数
2、状态转移方程：dp[i][j] = dp[i+1][j] + dp[i][j+1]

时间复杂度：O(m*n)
*/
func uniquePaths(m int, n int) int {
	// 二维数组的创建
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			right, down := 0, 0
			if j+1 < n {
				right = dp[i][j+1]
			}
			if i+1 < m {
				down = dp[i+1][j]
			}
			if i == m-1 && j == n-1 {
				continue
			}
			dp[i][j] = right + down
		}
	}
	fmt.Println(dp)
	return dp[0][0]
}

func main() {
	fmt.Println("vim-go")
	m, n := 3, 7
	result := uniquePaths(m, n)
	fmt.Println(result)
}
