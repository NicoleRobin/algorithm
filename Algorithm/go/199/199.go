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
func rightSideView(root *binarytree.TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	currentLevelNodes := []*binarytree.TreeNode{root}
	for len(currentLevelNodes) > 0 {
		result = append(result, currentLevelNodes[len(currentLevelNodes)-1].Val)
		var nextLevelNodes []*binarytree.TreeNode
		for _, node := range currentLevelNodes {
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		currentLevelNodes = nextLevelNodes
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	root := binarytree.BuildTree(nums)
	result := rightSideView(root)
	fmt.Println(result)
}
