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
func isValidBST(root *binarytree.TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(root *binarytree.TreeNode) (bool, int, int)
	dfs = func(root *binarytree.TreeNode) (bool, int, int) {
		if root == nil {
			return true, 0, 0
		}
		if root.Left == nil && root.Right == nil {
			return true, root.Val, root.Val
		}

		leftIsBst, rightIsBst := true, true
		max, min := root.Val, root.Val
		if root.Left != nil {
			left, leftMax, leftMin := dfs(root.Left)
			leftIsBst = left && (root.Val > leftMax)
			min = leftMin
		}
		if root.Right != nil {
			right, rightMax, rightMin := dfs(root.Right)
			rightIsBst = right && (root.Val < rightMin)
			max = rightMax
		}

		return leftIsBst && rightIsBst, max, min
	}
	isBst, _, _ := dfs(root)
	return isBst
}

func main() {
	fmt.Println("leetcode")
	nums := []int{5, 1, 4, -1, -1, 3, 6}
	root := binarytree.BuildTree(nums)
	fmt.Println(root)
	binarytree.LevelPrintTree(root)
	isBst := isValidBST(root)
	fmt.Println(isBst)
}
