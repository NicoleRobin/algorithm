package main

import "fmt"

/*

想办法利用有序的特点，使用二分查找

特殊case比较多，需注意：
1，2，3，4，5
2，1 这个是一开始没有想到的
3，4，5，1，2

思路：
1、如果有序数组经过了旋转，那么整个数组会形成两个上升子序列，且前面的

错误的，应该选择右端点进行比较，原因在下面的思路中有解释；
*/
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] <= nums[right] {
		return nums[left]
	}

	var result int
	for left < right {
		mid := (left + right) / 2
		if nums[left] == nums[mid] {
			result = nums[left]
			break
		} else if nums[left] < nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

/*
思路：旋转排序数组的特点：
1、旋转后的数组仍然是部分有序的，但是有一部分发生了断裂，形成两段单调递增的序列。
2、右端点 nums[right] 必然落在 最小值所在的那一段。
*/
func findMin2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

type TestCase struct {
	nums   []int
	result int
}

func main() {
	var testCases = [][]int{
		{3, 4, 5, 1, 2},
		{11, 13, 15, 17},
		{2, 1},
		{1, 2, 3, 4},
	}

	for _, testCase := range testCases {
		result := findMin2(testCase)
		fmt.Println("nums:", testCase, ", result:", result)
	}
}
