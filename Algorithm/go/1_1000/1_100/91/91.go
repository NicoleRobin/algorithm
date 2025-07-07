package main

import "fmt"

/*
思路：动态规划
1、定义状态：dp[i] 表示从i开始到末尾的解码方式
2、状态转移方程：
如果s[i:i+1]无法解码，则dp[i] = 0
否则：

	如果s[i:i+2]能解码
		如果s[i+1:i+2]能解码，则dp[i] = dp[i+2] + dp[i+1]
		否则dp[i] = dp[i+2]
	如果s[i:i+2]不能解码，则dp[i] = dp[i+1]

3、初始化：dp[n] = 1, dp[n-1] = decodeMap[s[n-1:n]]
5、返回结果：dp[0]

总结
1、还需要考虑不能因为前面的拆分导致后面的拆分不合法，例如2101，如果拆分为2 1 0 1，则01不合法，如果拆分为21 0 1，则01不合法，所以这里的0必须要1在一起，拆分为2 10 1；
2、上面的问题是由于前面的拆分可能会导致后面的拆分不合法，那是不是可以考虑从后往前考虑呢？
3、
*/
func numDecodings(s string) int {
	decodeMap := map[string]int{
		"1":  1,
		"2":  1,
		"3":  1,
		"4":  1,
		"5":  1,
		"6":  1,
		"7":  1,
		"8":  1,
		"9":  1,
		"10": 1,
		"11": 1,
		"12": 1,
		"13": 1,
		"14": 1,
		"15": 1,
		"16": 1,
		"17": 1,
		"18": 1,
		"19": 1,
		"20": 1,
		"21": 1,
		"22": 1,
		"23": 1,
		"24": 1,
		"25": 1,
		"26": 1,
	}

	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	dp[n-1] = decodeMap[s[n-1:n]]
	for i := n - 2; i >= 0; i-- {
		if decodeMap[s[i:i+1]] != 0 {
			if decodeMap[s[i:i+2]] != 0 {
				if decodeMap[s[i+1:i+2]] != 0 {
					dp[i] = dp[i+2] + dp[i+1]
				} else {
					dp[i] = dp[i+2]
				}
			} else {
				dp[i] = dp[i+1]
			}
		}
	}

	return dp[0]
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
	}
	for i, testCase := range testCases {
		ans := numDecodings(testCase.s)
		if ans != testCase.want {
			fmt.Printf("i: %d failed, ans: %v, want: %v\n", i, ans, testCase.want)
		} else {
			fmt.Printf("i: %d passed\n", i)
		}
	}
}
