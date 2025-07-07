package main

/*
思路：动态规划
1、定义状态：dp[i] 表示长度为i的good string的个数
2、状态转移方程：

	dp[i] = dp[i-zero] + dp[i-one]

3、初始化：dp[0] = 1
4、返回结果：dp[high] - dp[low-1]
*/
func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] += dp[i-zero]
		}
	}

	return dp[high] - dp[low-1]
}

func main() {}
