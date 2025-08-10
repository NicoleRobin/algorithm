package main

import (
	"fmt"
	"sort"
)

/*
思路：动态规划
固定一个维度，然后就可以将题转换为最长递增子序列，同时需要留意<=不同于<，=是不符合要求的；

1、状态定义
dp[i]: 第i个信封作为结尾的最大信封个数
2、状态转移方程
dp[i] = 0到i-1中，信封的宽和高都小于第i个信封的最大值

3、初始状态

4、结果
dp数组的最大值

时间复杂度：n*n
*/
func maxEnvelopes1(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return false
	})
	// fmt.Println(envelopes)

	n := len(envelopes)
	dp := make([]int, n)

	var ans int
	for i := 0; i < n; i++ {
		var maxValue int
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				if dp[j] > maxValue {
					maxValue = dp[j]
				}
			}
		}
		dp[i] = maxValue + 1

		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

/*
思路：二分查找

1、先排序，假设第一个值为宽，第二个值为高；按照宽升序，高降序排序；
2、然后遍历数组的高，如果当前的高大于已有的所有值，则附加到结尾，否则不做任何处理；
3、最终记录高的值的数组长度就是结果；

时间复杂度：n*logn
*/
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	var f []int
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}

func main() {
	testCases := []struct {
		envelopes [][]int
		want      int
	}{
		{
			envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
			want:      3,
		},
		{
			envelopes: [][]int{{1, 1}, {1, 1}, {1, 1}},
			want:      1,
		},
		{
			envelopes: [][]int{{1, 1}},
			want:      1,
		},
	}

	for i, testCase := range testCases {
		ans := maxEnvelopes(testCase.envelopes)
		if ans != testCase.want {
			fmt.Printf("Case #%d: failed, ans=%d, want %d\n", i+1, ans, testCase.want)
		} else {
			fmt.Printf("Case #%d: passed\n", i+1)
		}
	}
}
