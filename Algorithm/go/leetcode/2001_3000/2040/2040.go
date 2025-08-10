package main

import (
	"fmt"
	"sort"
)

/*
思考：
1、直观的想法是双层循环遍历两个数组，计算出所有的乘积，然后排序后取第k小的值。
但是这样子就没有利用到两个数组的有序性。

思路：暴力法
事件复杂度：O(n ^ 2)
*/
func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	var multiplyList []int64
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			multiplyList = append(multiplyList, int64(nums1[i])*int64(nums2[j]))
		}
	}
	sort.Slice(multiplyList, func(i, j int) bool {
		return multiplyList[i] < multiplyList[j]
	})
	fmt.Println(multiplyList)

	return multiplyList[k-1]
}

/*
思考：
1、暴力法的时间复杂度是O(n^2)，对于较大的数组会导致超时。
2、既然题目提到两个数组是有序地，那么就要想办法利用这个特性来进行优化。
3、但是数字是有正数和负数的，所以需要分开考虑。
4、将两个数组的正数和负数分开处理，记为pos1, pos2, neg1, neg2，所以从小到大排列的乘积有四种情况：
pos1 * neg2, pos2 * neg1, pos1 * pos2, neg1 * neg2
5、对于正数的情况，直接计算出所有的乘积，然后排序后取第k小的值。
6、把四个区域考虑为独立的四个区域，然后分别计算每个区域的乘积数量，在每个区域内等同于之前的在矩阵中查k小值的思路。

思路：二分查找

*/
func kthSmallestProduct2(nums1 []int, nums2 []int, k int64) int64 {
	var ans int64
	n, m, k := len(nums1), len(nums2), int(k)
	i0 := sort.SearchInts(nums1, 0) // 四个区域的水平分界线
	j0 := sort.SearchInts(nums2, 0) // 四个区域的垂直分界线

	return ans
}

func main() {
	testCases := []struct {
		nums1    []int
		nums2    []int
		k        int64
		expected int
	}{
		{
			nums1:    []int{-2, -1, 0, 1, 2},
			nums2:    []int{-3, -1, 2, 4, 5},
			k:        3,
			expected: -6,
		},
	}
	for i, tc := range testCases {
		result := kthSmallestProduct(tc.nums1, tc.nums2, tc.k)
		if result != int64(tc.expected) {
			println("Test case", i+1, "failed: expected", tc.expected, "but got", result)
		} else {
			println("Test case", i+1, "passed")
		}
	}
}
