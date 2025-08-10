package main

import "fmt"

/*
思路：
双指针，fast遍历旧数组，slow指针指向新数组的元素位置
*/

func removeElem(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}
	return slow
}

func main() {
	testCases := []struct {
		nums     []int
		val      int
		expected int
	}{
		{
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
	}

	for _, tt := range testCases {
		ret := removeElem(tt.nums, tt.val)
		if ret != tt.expected {
			fmt.Printf("nums:%+v, val:%d, expected:%d, got:%d\n", tt.nums, tt.val, tt.expected, ret)
		} else {
			fmt.Printf("nums:%+v, val:%d, ret:%d\n", tt.nums, tt.val, ret)
		}
	}
}
