package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
思路：
1、递归遍历二叉树
2、求每个节点左右两个子树的高度，经过该节点的路径即为左右子树高度相加在加上当前节点1
*/
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lHeight := dfs(root.Left)
		rHeight := dfs(root.Right)

		if lHeight+rHeight+1 > result {
			result = lHeight + rHeight + 1
		}

		return max(lHeight, rHeight) + 1
	}
	dfs(root)

	return result - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("vim-go")
}
