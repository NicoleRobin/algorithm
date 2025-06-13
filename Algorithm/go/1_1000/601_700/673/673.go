package main

/*
思路：动态规划
*/
func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))

	var maxLen, ans int
	for i, numI := range nums {
		dp[i] = 1
		cnt[i] = 1
		for j, numJ := range nums[:i] {
			if numI > numJ {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			ans = cnt[i]
		} else if dp[i] == maxLen {
			ans += cnt[i]
		}
	}
	return ans
}

func main() {

}
