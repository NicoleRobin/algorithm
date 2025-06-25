package main

/*
思路：动态规划
dp[i]表示以arr[i]结尾的最长等差子序列长度
时间复杂度：O(n^2)
*/
func longestSubsequence(arr []int, difference int) int {
	var ans int
	dp := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		dp[i] = 1 // 每个元素至少可以组成长度为1的等差子序列
		for j := 0; j < i; j++ {
			if arr[i]-arr[j] == difference {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}

	return ans
}

/*
思路：动态规划
dp[x]表示已x(arr[i])结尾的最长等差子序列长度
*/

func longestSubsequence1(arr []int, difference int) int {
	dp := make(map[int]int)
	var ans int
	for _, num := range arr {
		dp[num] = dp[num-difference] + 1
		if dp[num] > ans {
			ans = dp[num]
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		arr        []int
		difference int
		expected   int
	}{
		{
			arr:        []int{1, 2, 3, 4},
			difference: 1,
			expected:   4,
		},
	}
	for i, tc := range testCases {
		result := longestSubsequence(tc.arr, tc.difference)
		if result != tc.expected {
			println("Test case", i+1, "failed: expected", tc.expected, "but got", result)
		} else {
			println("Test case", i+1, "passed")
		}
	}
}
