package main

import "fmt"

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	var result int
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			result = mid
			break
		} else if nums[mid] < target {
			start = mid + 1
			result = start
		} else {
			end = mid - 1
			result = start
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 5, 6}
	result := searchInsert(nums, 5)
	fmt.Println(result)
}
