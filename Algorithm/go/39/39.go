package main

import (
	"fmt"
)

/*
思路：
1、回溯算法
2、递归函数 backtrack(start int, path []int, sum int)

和其他回溯算法的区别：
1、因为每个数字都可以重复使用，所以每次递归时，start 传入 i，而不是 i+1
*/
func combinationSum(candidates []int, target int) [][]int {
	var backtrack func(start int, path []int, sum int)

	var result [][]int
	backtrack = func(start int, path []int, sum int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}
		if sum >= target {
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(i, path, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(0, []int{}, 0)
	return result
}

func main() {
	fmt.Println("leetcode-39")
	candidates := []int{2, 7, 6, 3, 5, 1}
	target := 9
	result := combinationSum(candidates, target)
	fmt.Println(result)
}
