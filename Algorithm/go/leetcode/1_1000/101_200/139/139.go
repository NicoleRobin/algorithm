package main

import "fmt"

/*
动态规划
1、状态定义：dp[i]：以s[i]结尾的字符串是否在wordMap能找到
2、状态转移方程：dp[i] = j{0-i}(s[i-j]是否在WordMap中)

注意：
初始状态值的设置，然后在遍历中一定不要覆盖掉初始状态值
*/
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println("vim-go")
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	result := wordBreak(s, wordDict)
	fmt.Println(result)
}
