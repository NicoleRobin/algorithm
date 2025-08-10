package main

import "fmt"

/*
思路：
1、将两个数组合并成一个有序数组
2、找到中位数
*/
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	mergedNums := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			mergedNums = append(mergedNums, nums1[i])
			i++
		} else {
			mergedNums = append(mergedNums, nums2[j])
			j++
		}
	}
	if i < len(nums1) {
		mergedNums = append(mergedNums, nums1[i:]...)
	}
	if j < len(nums2) {
		mergedNums = append(mergedNums, nums2[j:]...)
	}
	fmt.Println("mergedNums:", mergedNums)

	var result float64
	if len(mergedNums)%2 == 0 {
		result = float64(mergedNums[len(mergedNums)/2]+mergedNums[len(mergedNums)/2-1]) / 2
	} else {
		result = float64(mergedNums[len(mergedNums)/2])
	}
	return result
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var result float64
	lenSum := len(nums1) + len(nums2)

	mid1, mid2 := 0, 0
	if lenSum%2 == 0 {

	} else {

	}

	return result
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}

	// result := findMedianSortedArrays1(nums1, nums2)
	result := findMedianSortedArrays2(nums1, nums2)
	// result := findMedianSortedArrays3(nums1, nums2)
	fmt.Println(result)
}
