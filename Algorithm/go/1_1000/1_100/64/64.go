package main

import "fmt"

/*
动态规划
1、状态定义dp[i][j]：从grid[i][j]到右下角的路径最小和
2、状态转移方程：dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]

时间复杂度：O(m*n)
空间复杂度：O(m*n)

这样子定义的状态方程计算起来比较复杂，还可以定义为dp[i][j]为从左上角到grid[i][j]节点的最短路径
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	xLen := len(grid)
	yLen := len(grid[0])
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	for i := xLen - 1; i >= 0; i-- {
		for j := yLen - 1; j >= 0; j-- {

			// 从右侧和下侧选择最小值，但是需要考虑无效值
			/*
				right, down := -1, -1
				if j+1 < yLen {
					right = dp[i][j+1]
				}
				if i+1 < xLen {
					down = dp[i+1][j]
				}
				dp[i][j] = grid[i][j]
				if right < 0 && down > 0 {
					dp[i][j] += down
				}
				if down < 0 && right > 0 {
					dp[i][j] += right
				}
				if right > 0 && down > 0 {
					if right > down {
						dp[i][j] += down
					} else {
						dp[i][j] += right
					}
				}
			*/

			// 根据不同情况进行处理
			if i == xLen-1 && j == yLen-1 {
				dp[i][j] = grid[i][j]
			} else if i == xLen-1 && j < yLen-1 {
				dp[i][j] = grid[i][j] + dp[i][j+1]
			} else if j == yLen-1 && i < xLen-1 {
				dp[i][j] = grid[i][j] + dp[i+1][j]
			} else {
				if dp[i+1][j] < dp[i][j+1] {
					dp[i][j] = grid[i][j] + dp[i+1][j]
				} else {
					dp[i][j] = grid[i][j] + dp[i][j+1]
				}
			}
		}
	}
	return dp[0][0]
}

/*
动态规划
1、状态定义dp[i][j]：从grid[0][0]到grid[i][j]的最小路径和
*/
func minPathSum2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	mLen := len(grid)
	nLen := len(grid[0])
	dp := make([][]int, mLen)
	for i := 0; i < mLen; i++ {
		dp[i] = make([]int, nLen)
	}
	// 边界状态
	dp[0][0] = grid[0][0]
	for i := 1; i < mLen; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < nLen; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < mLen; i++ {
		for j := 1; j < nLen; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[mLen-1][nLen-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("vim-go")
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	result := minPathSum(grid)
	fmt.Println(result)
}
