package main

/*
思路：动态规划
dp[i][j]: 表示以 nums[i] 结尾，公差为 j 的最长等差子序列的长度。
状态转移方程：
dp[i][j] = dp[k][j] + 1的最大值，其中 k < i 且 nums[i] - nums[k] = j。
*/
func longestArithSeqLength(nums []int) int {
	var ans int
	dp := make([]map[int]int, len(nums))
	dp[0] = make(map[int]int)
	for i := 1; i < len(nums); i++ {
		dp[i] = make(map[int]int)
		for j := i - 1; j >= 0; j-- {
			// for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			if dp[i][diff] == 0 {
				dp[i][diff] = dp[j][diff] + 1
				ans = max(ans, dp[i][diff])
			}
		}
	}

	return ans + 1
}

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 6, 9, 12},
			expected: 4,
		},
	}

	for i, tc := range testCases {
		result := longestArithSeqLength(tc.nums)
		if result != tc.expected {
			println("Test case", i+1, "failed: expected", tc.expected, "but got", result)
		} else {
			println("Test case", i+1, "passed")
		}
	}
}
