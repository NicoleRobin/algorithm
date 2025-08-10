package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
思考：
1、让1尽量靠左，但是怎么实现呢？
排序

思路：
暴力：
1、最多6种拼接方式，都组装出来比较一下；
优化方案：
1、排序，当遇到前缀相同的两个字符串时，比较两种拼接起来的大小
*/
func maxGoodNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		strI := fmt.Sprintf("%b", nums[i])
		strJ := fmt.Sprintf("%b", nums[j])
		return strI+strJ > strJ+strI
	})

	var strNums []string
	for _, num := range nums {
		strNums = append(strNums, fmt.Sprintf("%b", num))
	}
	res := strings.Join(strNums, "")
	bRes, err := strconv.ParseInt(res, 2, 64)
	if err != nil {
		return 0
	}
	return int(bRes)
}

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 3},
			expected: 30,
		},
	}
	for i, tc := range testCases {
		res := maxGoodNumber(tc.nums)
		if res != tc.expected {
			fmt.Printf("Case #%d fail, expected:%d, res:%d\n", i+1, tc.expected, res)
		} else {
			fmt.Printf("Case #%d pass, expected:%d, res:%d\n", i+1, tc.expected, res)
		}
	}
}
