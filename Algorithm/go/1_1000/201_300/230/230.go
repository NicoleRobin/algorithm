package main

import (
	"fmt"
	"github.com/nicolerobin/algorithm/go/binarytree"
)

/*
 * Definition for a binary tree node.
 */

func kthSmallest(root *binarytree.TreeNode, k int) int {
	var searchBST func(*binarytree.TreeNode)
	var result, count int
	searchBST = func(nRoot *binarytree.TreeNode) {
		if nRoot == nil {
			return
		}

		searchBST(nRoot.Left)

		count++
		if count == k {
			result = nRoot.Val
		}

		searchBST(nRoot.Right)
	}
	searchBST(root)
	return result
}

func main() {
	root := binarytree.BuildTree([]int{5, 3, 6, 2, 4, -1, -1, 1})
	binarytree.LevelPrintTree(root)
	result := kthSmallest(root, 3)
	fmt.Println("result:", result)
}
