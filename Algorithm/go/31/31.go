package main

import (
	"fmt"
)

/*
1、是否会重复
*/
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	// 从后向前，寻找第一个非降序对
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 从后向前寻找第一个大于num[i]的值
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 翻转从i+1到结尾的所有数字，因为i+1到结尾的所有数字是降序的，所以反转之后变成升序了
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func main() {
	fmt.Println("leetcode-31")
	nums := []int{6, 5, 4, 6, 7, 1}
	fmt.Printf("nums:%+v\n", nums)
	nextPermutation(nums)
	fmt.Printf("nums:%+v\n", nums)
}
