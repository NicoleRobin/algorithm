package main

/*
思路：动态规划
dp[i][j]: 表示以 nums[i] 结尾，公差为 j 的最长等差子序列的长度。
*/
func longestArithSeqLength(nums []int) int {
	var ans int

	return ans
}

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 3},
			expected: 3,
		},
	}

	for i, tc := range testCases {
		result := longestArithSeqLength(tc.nums)
		if result != tc.expected {
			panic("test case " + string(i) + " failed: expected " + string(tc.expected) + ", got " + string(result))
		}
	}
}
