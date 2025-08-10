package main

import (
	"fmt"
	"slices"
	"strconv"
)

/*
思路：回溯

总结：
1、时间复杂度太高
*/
func reorderedPowerOf2(n int) bool {
	strN := fmt.Sprintf("%d", n)
	lenStrN := len(strN)

	var res bool
	var backtrack func([]byte)
	selectedIndex := map[int]bool{}
	backtrack = func(memo []byte) {
		if len(memo) == lenStrN {
			if len(memo) > 0 && memo[0] == '0' {
				return
			}
			iMemo, _ := strconv.ParseInt(string(memo), 10, 64)
			if isPowerOfTwo(iMemo) {
				res = true
				return
			}
		}
		for i := 0; i < len(strN); i++ {
			if selectedIndex[i] {
				continue
			}
			selectedIndex[i] = true
			memo = append(memo, strN[i])
			backtrack(memo)
			delete(selectedIndex, i)
			memo = memo[:len(memo)-1]
		}
	}
	backtrack([]byte{})

	return res
}

func isPowerOfTwo(n int64) bool {
	return n&(n-1) == 0
}

/*
思路：预处理+hash表
1、先将所有的2的幂记录下来，然后将n转换为排序的字符串并判断是否存在于hash表；
*/
func reorderedPowerOf2_1(n int) bool {
	powTwoSortedStrSet := map[string]bool{}
	for i := 1; i < 1_000_000_000; i <<= 1 {
		s := intToSortStr(i)
		powTwoSortedStrSet[s] = true
	}

	s := intToSortStr(n)
	return powTwoSortedStrSet[s]
}

func intToSortStr(n int) string {
	s := []byte(strconv.Itoa(n))
	slices.Sort(s)
	return string(s)
}

func main() {
	testCases := []struct {
		n        int
		expected bool
	}{
		{
			n:        10,
			expected: false,
		},
	}
	for i, tc := range testCases {
		res := reorderedPowerOf2(tc.n)
		if res != tc.expected {
			fmt.Printf("Case: %d fail, Expected: %v, Result: %v\n", i, tc.expected, res)
		} else {
			fmt.Printf("Case: %d pass, Result: %v\n", i, tc.expected)
		}
	}
}
