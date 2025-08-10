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
func dfs(root *binarytree.TreeNode, targetSum int, path []int, result *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		path = append(path, root.Val)
		tmpPath := make([]int, len(path))
		copy(tmpPath, path)
		*result = append(*result, tmpPath)
		return
	}

	path = append(path, root.Val)
	if root.Left != nil {
		dfs(root.Left, targetSum-root.Val, path, result)
	}
	if root.Right != nil {
		dfs(root.Right, targetSum-root.Val, path, result)
	}
}

func pathSum(root *binarytree.TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	dfs(root, targetSum, []int{}, &result)
	return result
}

func main() {
	nums := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}
	root := binarytree.BuildTree(nums)
	binarytree.LevelPrintTree(root)
	paths := pathSum(root, 22)
	fmt.Println(paths)
}
