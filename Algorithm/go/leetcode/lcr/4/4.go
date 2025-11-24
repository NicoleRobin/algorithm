package main

import "fmt"

/*
思考：
1、需要考虑负数的情况；
*/
func singleNumber(nums []int) int {
	var ans int

	for bit := 0; bit < 64; bit++ {
		var oneCount int
		for _, num := range nums {
			if num>>bit&1 == 1 {
				oneCount++
			}
		}
		if oneCount%3 == 1 {
			ans |= 1 << bit
		}
	}

	return ans
}

func main() {
	testCases := [...]struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 2, -3, 2},
			expected: -3,
		},
	}
	for i, tc := range testCases {
		result := singleNumber(tc.nums)
		if result != tc.expected {
			fmt.Printf("Case %d fail, Expected %d, got %d\n", i, tc.expected, result)
		} else {
			fmt.Printf("Case %d Pass\n", i)
		}
	}
}
