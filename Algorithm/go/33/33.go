package main

import "fmt"

/*
思考：
1、有序数组经过旋转后，仍然是部分有序的，但是有一部分发生了断裂，形成了两个单调递增的序列。
2、仍然想办法利用二分搜索，只是在判断的时候，需要考虑到旋转的情况。

思路：
1、二分搜索，比较中点值和目标值，如果相等，则返回。
2、如果目标值小于中点值，则需要向数值更小的范围搜索，否则向数值更大的范围搜索。

注意：
1、left <= right 这个条件是必须的，因为需要考虑到只有一个元素的情况。相比于153题，这里的条件是 <=
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var result = -1
	for left <= right {
		fmt.Println("left:", left, ", right:", right)
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			break
		} else if target > nums[mid] {
			// 往更大的方向找
			if nums[mid] < nums[right] {
				if target > nums[right] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				left = mid + 1
			}
		} else {
			// 往更小的方向找
			if nums[mid] < nums[right] {
				right = mid - 1
			} else {
				if target < nums[left] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}
	return result
}

type TestCase struct {
	nums   []int
	result int
}

func main() {
	testCases := []TestCase{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
		{[]int{1}, 0},
	}

	for _, testCase := range testCases {
		result := search(testCase.nums, testCase.result)
		fmt.Println("testCase:", testCase.nums, ", result:", result)
	}
}
