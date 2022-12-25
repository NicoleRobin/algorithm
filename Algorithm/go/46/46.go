package main

import "fmt"

/*
求全排列，回溯法
和组合的区别在于递归的每一层并非从start开始，而是从0开始，
只是需要通过used辅助数组记录已经使用过的数字，避免重复
*/
func permute(nums []int) [][]int {
	var backtracking func(used map[int]bool, path []int)
	res := [][]int{}
	backtracking = func(used map[int]bool, path []int) {
		if len(path) >= len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if _, ok := used[nums[i]]; ok {
				continue
			}
			used[nums[i]] = true
			path = append(path, nums[i])
			backtracking(used, path)
			path = path[:len(path)-1]
			delete(used, nums[i])
		}
	}
	used := map[int]bool{}
	backtracking(used, []int{})
	return res
}

func main() {
	testCases := []struct {
		nums []int
	}{
		{
			nums: []int{1, 2, 3},
		},
	}

	for _, tt := range testCases {
		res := permute(tt.nums)
		fmt.Printf("nums:%+v, res:%+v\n", tt.nums, res)
	}
}
