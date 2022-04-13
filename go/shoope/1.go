/*
给定一个只包含小写字母的字符串，允许您用任何字母替换不超过 k 个字母，请找出替换后具有相同字母的最长子串的长度

详细描述

示例 1：
输入：str = "aabccbb"，k=2
输出：5
解释：将两个 'c' 替换为 'b' 以获得最长的重复子串 "bbbbb"

示例 2：
输入：str = "abbcb", k=1
输出：4
说明：将 'c' 替换为 'b' 以获得最长的重复子字符串 "bbbb"

示例 3：
输入：str = "abccde"，k=1
输出：3
说明：将 'b' 或 'd' 替换为 'c' 以获得最长的重复子字符串 "ccc"
*/
package main

import (
	"fmt"
	"math"
)

func longestStr2(str string, k int) int {
	left, right := 0, 1
	max := math.MinInt32
	count := 1
	replace := k
	notEqual := -1
	for right < len(str) {
		// fmt.Println(right)
		if str[right] == str[left] { // 这种判断方法不对，例如abbb，k=1时，最长长度应该是4的，但是算出来是3，因为总是以第一个单词为锚点
			count++
			right++
		} else {
			if notEqual == -1 {
				notEqual = right
			}
			if replace > 0 {
				replace--
				count++
				right++
			} else {
				replace = k
				count = 1

				left = notEqual
				right = notEqual + 1
				notEqual = -1
			}
		}
		if count > max {
			max = count
		}
		fmt.Printf("right:%d, count:%d\n", right, count)
	}
	return max
}

func longestStr(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	state := make([]int, len(str))
	state[0] = 1
	replace := k
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			state[i] = state[i-1] + 1
		} else {
			if replace > 0 {
				replace--
				state[i] = state[i-1] + 1
			} else {
				state[i] = 1
				replace = k
			}
		}
	}
	fmt.Println(state)
	max := math.MinInt32
	for i := 0; i < len(state); i++ {
		if state[i] > max {
			max = state[i]
		}
	}

	return max
}

type TestCase struct {
	Str string
	Len int
}

func main() {
	testCases := []TestCase{
		// {"aabccbb", 2},
		// {"abbcb", 1},
		// {"abccde", 1},
		// {"aababba", 1},
		{"abbb", 2},
	}
	for _, testCase := range testCases {
		result := longestStr2(testCase.Str, testCase.Len)
		fmt.Println(result)
	}

}
