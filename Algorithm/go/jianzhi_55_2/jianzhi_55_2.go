package main

import (
	"fmt"

	"github.com/nicolerobin/algorithm/go/binarytree"
)

/*
断一个二叉树是否是平衡二叉树
*/

func isBalanced2(root *binarytree.TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHeight, left := isBalanced2(root.Left)
	rightHeight, right := isBalanced2(root.Right)

	height := leftHeight + 1
	if rightHeight > leftHeight {
		height = rightHeight + 1
	}

	if leftHeight-rightHeight >= -1 && leftHeight-rightHeight <= 1 {
		fmt.Printf("height:%d, balanced:%t\n", height, left && right)
		return height, left && right
	} else {
		fmt.Printf("height:%d, balanced:%t\n", height, false)
		return height, false
	}
}

func isBalanced(root *binarytree.TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight, left := isBalanced2(root.Left)
	rightHeight, right := isBalanced2(root.Right)

	if leftHeight-rightHeight >= -1 && leftHeight-rightHeight <= 1 {
		return left && right
	} else {
		return false
	}
}

func main() {
	fmt.Println("vim-go")
	// nums := []int{3, 9, 20, -1, -1, 15, 7}
	nums := []int{1, 2, 2, 3, 3, -1, -1, 4, 4}
	root := binarytree.BuildTree(nums)
	binarytree.PrePrintTree(root)
	isBalanced := isBalanced(root)
	fmt.Println("isBalanced:", isBalanced)
}
