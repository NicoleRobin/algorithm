package main

import (
	"fmt"
	"sort"
)

type NumCount struct {
	Num   int
	Count int
}

/*
思路：
1、将数组转换为num-count的map；
2、将map转换为元素为包含num和count的结构体的数组；
3、按照count倒序排序；
4、取前top k的元素；

时间复杂度：O(nlogn)
空间复杂度：O(n)
*/
func topKFrequent(nums []int, k int) []int {
	numCountMap := make(map[int]int)
	for _, num := range nums {
		numCountMap[num]++
	}

	var numCountList []NumCount
	for num, count := range numCountMap {
		numCountList = append(numCountList, NumCount{
			Num:   num,
			Count: count,
		})
	}

	sort.Slice(numCountList, func(i, j int) bool {
		return numCountList[i].Count > numCountList[j].Count
	})

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, numCountList[i].Num)
	}
	return result
}

func main() {
	fmt.Println("leetcode-347")
}
