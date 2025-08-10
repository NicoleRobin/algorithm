package main

import (
	"math/bits"
	"strconv"
	"strings"
)

/*
思考：
1、子序列，可以跳着选择
*/
func longestSubsequence(s string, k int) int {
	n, m := len(s), bits.Len(uint(k))
	if n < m {
		return n
	}

	ans := m
	sufVal, _ := strconv.ParseInt(s[n-m:], 2, 0)
	if int(sufVal) > k {
		ans--
	}
	return ans + strings.Count(s[:n-m], "0") // 添加前导零
}

func strBinaryToInt(s string) int64 {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0
	}
	return num
}

func main() {
	testCases := []struct {
		s    string
		k    int
		want int
	}{
		{
			s:    "1001010",
			k:    5,
			want: 5,
		},
		{
			s:    "0010101",
			k:    1,
			want: 6,
		},
	}
	for i, tc := range testCases {
		got := longestSubsequence(tc.s, tc.k)
		if got != tc.want {
			println("Test case", i, "failed: got", got, "want", tc.want)
		} else {
			println("Test case", i, "passed")
		}
	}
}
