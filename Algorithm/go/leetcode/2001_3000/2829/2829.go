package main

import "fmt"

/*
思路：贪心算法
1、从1开始，每次加上start，如果k-start存在，则start++
*/
func minimumSum(n int, k int) int {
	var result int
	numSet := make(map[int]bool)
	var count int
	start := 1
	for count < n {
		if _, ok := numSet[k-start]; ok {
			start++
		} else {
			result += start
			numSet[start] = true
			start++
			count++
		}
	}
	return result
}

func main() {
	result := minimumSum(5, 2)
	fmt.Println(result)
}
