package main

import (
	"fmt"
	"runtime/pprof"
	"sort"
)

func minimizeMax(nums []int, p int) int {
}
func main() {
	testCases := []struct {
		nums     []int
		p        int
		expected int
	}{
		{
			nums:     []int{},
			p:        0,
			expected: 0,
		},
	}

	for i, testCase := range testCases {
		result := minimizeMax(testCase.nums, testCase.p)
		fmt.Println("i:", i, "result:", result, "expected:", testCase.expected)
	}
}
