package main

import "fmt"

/*
思考：
1、要求的数对是相等的，该怎么快速找到相等的呢？
2、
思路：
1、
*/
func countPairs(nums []int, k int) int {
	var numCountMap = make(map[int][]int)
	var result int
	for i, num := range nums {
		if _, ok := numCountMap[num]; ok {
			for _, index := range numCountMap[num] {
				if isDivide(i*index, k) {
					result++
				}
			}
			numCountMap[num] = append(numCountMap[num], i)
		} else {
			numCountMap[num] = []int{i}
		}
	}

	return result
}

func isDivide(x, y int) bool {
	divide := x / y
	return divide*y == x
}

func main() {
	nums := []int{3, 1, 2, 2, 2, 1, 3}
	k := 2
	result := countPairs(nums, k)
	fmt.Println("result:", result)
}
