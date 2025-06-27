package main

/*
思路：动态规划
dp[i][j]: 表示使用i个0和j个1可以形成的最大子集的大小
状态转移方程：
*/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

}
func main() {

}
