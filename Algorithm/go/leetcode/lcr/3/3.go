package main

import (
	"fmt"
	"math/bits"
	"reflect"
)

func countBits(n int) []int {
	var ans []int
	for i := 0; i <= n; i++ {
		// ans = append(ans, countBit(i))
		ans = append(ans, countBit1(i))
	}
	return ans
}

func countBit1(n int) int {
	return bits.OnesCount(uint(n))
}

/*
Brian Kernighan 算法
*/
func countBit(n int) int {
	count := 0
	for n != 0 {
		n &= n - 1 // 清除最低位的 1
		count++
	}
	return count
}

func main() {
	testCases := []struct {
		n        int
		expected []int
	}{
		{
			n:        2,
			expected: []int{0, 1, 1},
		},
	}
	for _, tc := range testCases {
		ans := countBits(tc.n)
		if !reflect.DeepEqual(ans, tc.expected) {
			fmt.Printf("\nExpected: %d, Actual: %d\n", tc.expected, ans)
		} else {
			fmt.Printf("\nNo. %d is %d\n", tc.n, ans)
		}
	}
}
