package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	var backtrack func([]int, int, int)
	backtrack = func(path []int, start, target int) {
		if target == 0 {
			fmt.Println(path)
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		if target < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] {
				continue
			}
			fmt.Printf("start:%d, path:%+v\n", start, path)
			path = append(path, candidates[i])
			backtrack(path, i+1, target-candidates[i])
			fmt.Printf("start:%d, path:%+v\n", start, path)
			path = path[:len(path)-1]
		}
	}

	path := []int{}
	backtrack(path, 0, target)
	return result
}

func main() {
	fmt.Println("leetcode-40")
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	result := combinationSum2(candidates, target)
	fmt.Println(result)
}
