package main

/*
思路：
1、实现一个函数计算一个全1的字符串中1的子串个数
*/

const (
	maxValue = 1000000000 + 7
)

func numSub(s string) int {
	var ans int
	start := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '1' && start == -1 {
			start = i
		}

		if s[i] == '0' && start != -1 {
			// count sub string and reset start
			count := subCount(s[start:i])
			start = -1
			ans += count
			ans = ans % maxValue
		}
	}
	if start != -1 {
		count := subCount(s[start:len(s)])
		start = -1
		ans += count
		ans = ans % maxValue
	}

	return ans
}

func subCount(s string) int {
	n := len(s)
	var ans int
	for i := 1; i <= n; i++ {
		ans += i
	}
	return ans
}

func main() {}
