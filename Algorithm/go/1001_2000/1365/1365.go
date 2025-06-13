package main

import "slices"

/*
思路：更高效的方法是计数排序
*/
func smallerNumbersThanCurrent(nums []int) []int {
	numIndexMap := map[int][]int{}
	for i, num := range nums {
		if _, ok := numIndexMap[num]; ok {
			numIndexMap[num] = append(numIndexMap[num], i)
		} else {
			numIndexMap[num] = []int{i}
		}
	}

	slices.Sort(nums)
	result := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			for _, index := range numIndexMap[num] {
				result[index] = i
			}
		} else {
			if num == nums[i-1] {
				continue
			} else {
				for _, index := range numIndexMap[num] {
					result[index] = i
				}
			}
		}
	}
	return result
}

func main() {

}
