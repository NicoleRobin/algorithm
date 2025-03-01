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
func lowestCommonAncestor(root, p, q *binarytree.TreeNode) *binarytree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func main() {
	root := binarytree.BuildTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})
	binarytree.LevelPrintTree(root)
	parentNode := lowestCommonAncestor(root, root.Left, root.Right)
	fmt.Println(parentNode.Val)
}
