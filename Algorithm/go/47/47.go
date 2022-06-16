package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	backtrack(nums, 0, &result)
	return result
}

func backtrack(nums []int, index int, result *[][]int) {
	fmt.Printf("nums:%+v, index:%d, result:%+v\n", nums, index, result)
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}

	numMap := map[int]interface{}{}
	for i := index; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; ok {
			continue
		} else {
			numMap[nums[i]] = struct{}{}
		}
		nums[index], nums[i] = nums[i], nums[index]
		fmt.Printf("index:%d, before:%+v\n", index, nums)
		backtrack(nums, index+1, result)
		nums[i], nums[index] = nums[index], nums[i]
		fmt.Printf("index:%d, after:%+v\n", index, nums)
	}
}

func main() {
	fmt.Println("leetcode-47")
	nums := []int{1, 1, 2}
	allList := permuteUnique(nums)
	fmt.Println(allList)
}
