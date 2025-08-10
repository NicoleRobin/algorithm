package main

import "sort"

/*
思考：
1、先排序，对整个数对数组排序；但是按第一个排序还是按第二个排序呢？
2、

思路：动态规划
dp[i]：表示以第i个数对为结尾的最长链长度。
状态转移方程：
在从0到i-1之间选择第二个数字小于当前第一个数对中dp[j]值最大的；

反思：
1、动态规划中，为什么有的状态定义是以第i个值为结尾，而有的状态定义是
*/
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	dp := make([]int, len(pairs))
	for i := 0; i < len(pairs); i++ {
		dp[i] = 1 // 每个数对至少可以形成一个链
		for j, q := range pairs {
			if pairs[i][0] > q[1] { // 如果当前数对的第一个元素大于前一个数对的第二个元素
				dp[i] = max(dp[i], dp[j]+1) // 更新最长链长度
			}
		}
	}

	return dp[len(pairs)-1]
}

/*
思路：贪心算法

*/
func findLongestChain1(pairs [][]int) int {
}

func main() {
}
