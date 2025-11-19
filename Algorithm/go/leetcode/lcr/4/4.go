package main

import "fmt"

func singleNumber(nums []int) int {
	var ans int
	for bit := 0; bit < 32; bit++ {
		var oneCount int
		for _, num := range nums {
			fmt.Printf("bit:%d, num:%b, test:%b\n", bit, num, 1<<bit)
			if num&1<<bit == 1 {
				oneCount++
			}
		}
		if oneCount%3 == 1 {
			ans |= 1 << bit
		}
		fmt.Printf("bit: %d, oneCount: %d\n", bit, oneCount)
	}

	return ans
}

func main() {
	testCases := [...]struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 2, 3, 2},
			expected: 3,
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
