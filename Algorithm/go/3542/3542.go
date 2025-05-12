package main

import (
	"fmt"
	"sort"
)

func minOperations(nums []int) int {
	zeroIndexMap := make(map[int]bool)
	numIndexListMap := make(map[int][]int)
	var uniqNumList []int
	for i, num := range nums {
		if num == 0 {
			zeroIndexMap[i] = true
		} else {
			if _, ok := numIndexListMap[num]; !ok {
				numIndexListMap[num] = []int{i}
				uniqNumList = append(uniqNumList, num)
			} else {
				numIndexListMap[num] = append(numIndexListMap[num], i)
			}
		}
	}

	var res int
	sort.Ints(uniqNumList)
	// fmt.Println("uniqNumList:", uniqNumList)
	for _, num := range uniqNumList {
		numIndexList := numIndexListMap[num]
		// fmt.Println("numIndexList:", numIndexList)
		// fmt.Println("zeroIndexMap:", zeroIndexMap)
		if len(numIndexList) == 1 {
			res++
			zeroIndexMap[numIndexList[0]] = true
		} else {
			res++
			zeroIndexMap[numIndexList[0]] = true
			for j := 1; j < len(numIndexList); j++ {
				zeroIndexMap[numIndexList[j]] = true
				for k := numIndexList[j-1] + 1; k < numIndexList[j]; k++ {
					if _, ok := zeroIndexMap[k]; ok {
						res++
						break
					}
				}
			}
		}
	}

	return res
}

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 1, 2, 1, 2},
			expected: 4,
		},
	}

	for i, testCase := range testCases {
		res := minOperations(testCase.nums)
		fmt.Println("i:", i, "res:", res, "expected:", testCase.expected)
	}
}
