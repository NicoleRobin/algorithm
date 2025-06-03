package main

import "golang.org/x/text/feature/plural"

/*
思路：动态规划-01背包问题
dp[i][j]表示：在最多i个0和j个1的情况下，最多能够选多少个字符串

*/
func findMaxForm(strs []string, m int, n int) int {
	zeroCountList, oneCountList := make([]int, len(strs)), make([]int, len(strs))
	for i, str := range strs {
		for _, ch := range str {
			if ch == '0' {
				zeroCountList[i]++
			} else if ch == '1' {
				oneCountList[i]++
			}
		}
	}

	var ans int
	for i, str := range strs {

	}

	return ans
}

func main() {
}
