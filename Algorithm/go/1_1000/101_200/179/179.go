package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
思考：
1、让较大的数字尽量靠左，但是怎么实现呢？
2、可以考虑排序，但是有特殊情况，如果354、和35比较，该选哪个呢？
将两个数拼接在一起，选择拼接结果更大的一种方式
35435 < 35354，所以选择将35放在前面
3、

思路：
1、
*/
func largestNumber(nums []int) string {
	var strNums []string
	for _, num := range nums {
		strNums = append(strNums, strconv.Itoa(num))
	}

	sort.Slice(strNums, func(i, j int) bool {
		strIJ := strNums[i] + strNums[j]
		strJI := strNums[j] + strNums[i]
		return strIJ > strJI
	})

	res := strings.Join(strNums, "")
	// 删除前缀0
	res = strings.TrimLeft(res, "0")
	if len(res) == 0 {
		return "0"
	}
	return res
}

func main() {
	testCases := []struct {
		nums     []int
		expected string
	}{
		{
			nums:     []int{0, 0},
			expected: "0",
		},
	}
	for i, tc := range testCases {
		res := largestNumber(tc.nums)
		if res != tc.expected {
			fmt.Printf("Case #%d fail, expected:%s, res:%s\n", i+1, tc.expected, res)
		} else {
			fmt.Printf("Case #%d pass, expected:%s, res:%s\n", i+1, tc.expected, res)
		}
	}
}
