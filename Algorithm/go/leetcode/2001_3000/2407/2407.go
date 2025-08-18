package main

/*
思路：动态规划
1、但是超时了，难怪是hard题，还得继续想办法优化，空间换时间？

在求解「上升子序列（IS）」问题时，一般有两种优化方法：
1、维护固定长度的 IS 的末尾元素的最小值 + 二分优化；
2、基于值域的线段树、平衡树等数据结构优化。
*/
func lengthOfLIS(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		maxLen := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && nums[i]-nums[j] <= k {
				if dp[j]+1 > maxLen {
					maxLen = dp[j] + 1
				}
			}
		}
		dp[i] = maxLen
		if maxLen > res {
			res = maxLen
		}
	}
	return res
}

func main() {

}
