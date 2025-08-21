package main

import "fmt"

/*
思路：动态规划
但是好像是错的
*/
func numSubmat1(mat [][]int) int {
	if len(mat) == 0 {
		return 0
	}

	m := len(mat)
	n := len(mat[0])
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				if dp[i+1][j] == 0 && dp[i][j+1] == 0 {
					dp[i+1][j+1] = 1
				} else if dp[i+1][j] != 0 && dp[i][j+1] != 0 {
					dp[i+1][j+1] = dp[i+1][j] + dp[i][j+1] - dp[i][j] + 1
				} else {
					if dp[i][j+1] != 0 {
						dp[i+1][j+1] = dp[i][j+1] - dp[i][j] + 1
						if j-1 >= 0 {
							dp[i+1][j+1] += dp[i][j-1]
						}
					} else {
						dp[i+1][j+1] = dp[i+1][j] - dp[i][j] + 1
						if i-1 >= 0 {
							dp[i+1][j+1] += dp[i-1][j]
						}
					}
				}
				ans += dp[i+1][j+1]
			}
		}
	}
	fmt.Println(dp)
	return ans
}

/*
思路：枚举矩形的上下边界
时间复杂度 O(m^2 * n)
*/
func numSubmat(mat [][]int) int {
	if len(mat) == 0 {
		return 0
	}
	m := len(mat)
	n := len(mat[0])

	var ans int
	for top := 0; top < m; top++ { // 枚举上边界
		a := make([]int, n)                       // 记录每一列从上到下连续1的个数
		for bottom := top; bottom < m; bottom++ { // 枚举下边
			h := bottom - top + 1 // 矩形的高度
			// 2348 全 h 子数组的数目
			last := -1
			for j := 0; j < n; j++ { // 枚举每一列
				a[j] += mat[bottom][j] // 更新当前列的连续1的个数
				if a[j] != h {         // 如果当前列的连续1的个数大
					last = j
				} else {
					ans += j - last // 计算当前列的连续1的个数为 h 的子数组的数目
				}
			}
		}
	}
	return ans
}

/*
思路：单调栈
前置题目 85
*/
func numSubmat2(mat [][]int) int {
	if len(mat) == 0 {
		reutrn 0
	}
	m := len(mat)
	n := len(mat[0])

	var ans int

	return ans
}

func main() {
	testCases := []struct {
		mat      [][]int
		expected int
	}{
		{
			mat:      [][]int{{0, 0, 0}, {0, 0, 0}, {0, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			expected: 12,
		},
		{
			// mat:      [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}},
			mat:      [][]int{},
			expected: 13,
		},
		{
			// mat:      [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}},
			mat:      [][]int{},
			expected: 5,
		},
	}
	for i, tc := range testCases {
		res := numSubmat(tc.mat)
		if res != tc.expected {
			println("Case #", i, " fail, expected ", tc.expected, " got ", res)
		} else {
			println("Case #", i, " pass")
		}
	}
}
