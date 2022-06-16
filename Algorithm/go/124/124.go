package main

import (
	"fmt"

	"github.com/nicolerobin/algorithm/go/binarytree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	max := a[0]
	for _, num := range a {
		if num > max {
			max = num
		}
	}
	return max
}

func maxPathSum(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(root *binarytree.TreeNode) int
	maxSum := root.Val
	dfs = func(root *binarytree.TreeNode) int {
		if root == nil {
			return 0
		}

		leftSum := dfs(root.Left)
		rightSum := dfs(root.Right)
		returnSum := max([]int{root.Val, root.Val + leftSum, root.Val + rightSum})
		sum := max([]int{root.Val, root.Val + leftSum, root.Val + rightSum, root.Val + leftSum + rightSum})
		if sum > maxSum {
			maxSum = sum
		}
		return returnSum
	}
	dfs(root)

	return maxSum
}

func main() {
	fmt.Println("leetcode")
	nums := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}
	root := binarytree.BuildTree(nums)
	binarytree.LevelPrintTree(root)

	maxSum := maxPathSum(root)
	fmt.Println(maxSum)
}
