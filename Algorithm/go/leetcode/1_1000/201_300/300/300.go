package main

/*
dp[i]: 以nums[i]结尾的最长上升子序列的长度
时间复杂度： O(n^2)
*/
func lengthOfLIS1(nums []int) int {
	dp := make([]int, len(nums))
	var maxVal int
	for i, num := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < num {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxVal = max(maxVal, dp[i])
	}
	return maxVal
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	var maxVal int
	for i, num := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < num {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxVal = max(maxVal, dp[i])
	}
	return maxVal
}

func main() {}
