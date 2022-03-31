package main

import (
	"fmt"

	"github.com/nicolerobin/algorithm/go/binarytree"
)

func sortedArrayToBST(nums []int) *binarytree.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	middle := len(nums) / 2
	var left *binarytree.TreeNode = nil
	if middle > 0 {
		left = sortedArrayToBST(nums[0:middle])
	}
	var right *binarytree.TreeNode = nil
	if middle+1 <= len(nums) {
		right = sortedArrayToBST(nums[middle+1:])
	}
	root := &binarytree.TreeNode{
		Val:   nums[middle],
		Left:  left,
		Right: right,
	}

	return root
}

func main() {
	fmt.Println("vim-go")
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	binarytree.LevelPrintTree(root)
}
