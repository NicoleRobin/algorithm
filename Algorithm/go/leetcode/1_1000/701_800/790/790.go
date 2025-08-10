package main

/*
思路：动态规划
1、状态定义，dp[i]: 长度为i的平铺数量
2、状态转移方程：动态规划分为两种：枚举、选或者不选，这里考虑使用枚举
dp[i] = dp[i-1] + d[i-2] + 2 + dp[i-3] + 2
3、初始状态
dp[0] = 0
dp[1] = 1
dp[2] = 2
4、结果为dp[n]
*/
func numTilings(n int) int {
	if n < 3 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3] + 2
	}
	return dp[n]
}

func numTilings1(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1]*2 + dp[i-3]) % (1e9 + 7)
	}
	return dp[n]
}

func main() {

}
