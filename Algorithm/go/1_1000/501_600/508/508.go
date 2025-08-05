package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode) int
	subTreeMap := make(map[int]int)
	dfs = func(root *TreeNode) int {
		fmt.Println(root)
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		subTreeMap[left+right+root.Val]++
		return left + right + root.Val
	}
	dfs(root)

	var maxCount int
	for _, count := range subTreeMap {
		if count > maxCount {
			maxCount = count
		}
	}
	for subTreeSum, count := range subTreeMap {
		if count == maxCount {
			res = append(res, subTreeSum)
		}
	}

	return res
}

func main() {
	testCases := []struct {
		root     *TreeNode
		expected []int
	}{
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: -3,
				},
			},
			expected: []int{2, -3, 4},
		},
	}
	for i, tc := range testCases {
		res := findFrequentTreeSum(tc.root)
		if !reflect.DeepEqual(res, tc.expected) {
			fmt.Printf("#%d failed: got %v, expected %v\n", i, res, tc.expected)
		} else {
			fmt.Printf("#%d pass: got %v, expected %v\n", i, res, tc.expected)
		}
	}
}
