package main

import "fmt"

/*
动态规划
1、状态定义：dp[i]：表示以下标i字符结尾的最长有效括号的长度
2、状态转移方程：
当前为(时dp[i] = 0
当前为)时，如果s[i-1]为(，则dp[i] = dp[i-2]+2；当s[i-dp[i-1]-1]为(时，dp[i]=dp[i-1]+2+dp[i-dp[i-1]-2]
*/
func longestValidParentheses(s string) int {
	result := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1] > 2 {
						dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}

		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func main() {
	fmt.Println("vim-go")
	s := ")()())"
	result := longestValidParentheses(s)
	fmt.Println(result)
}
