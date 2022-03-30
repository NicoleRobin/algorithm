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
func topK(root *binarytree.TreeNode, k *int) *binarytree.TreeNode {
	if root == nil {
		return nil
	}

	right := topK(root.Right, k)
	if right != nil {
		return right
	}

	*k--
	if *k == 0 {
		return root
	}

	left := topK(root.Left, k)
	if left != nil {
		return left
	}

	return nil
}

func kthLargest(root *binarytree.TreeNode, k int) int {
	topKNode := topK(root, &k)
	return topKNode.Val
}

func main() {
	fmt.Println("vim-go")
	nums := []int{3, 1, 4, -1, 2}
	root := binarytree.BuildTree(nums)
	topK := kthLargest(root, 1)
	fmt.Printf("topK:%d, value:%d\n", 1, topK)
}
